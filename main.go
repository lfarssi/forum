package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strings"

	"forum/app/models"
	"forum/routes"
	"net/http"
)

func LoadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || !strings.Contains(line, "=") {
			continue // Skip comments and invalid lines
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			os.Setenv(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		}
	}
	return scanner.Err()
}
const dbPassword = "jononl3adama"

func authenticate(password string) bool {
	return password == dbPassword
}
func main() {
	var inputPassword string
	fmt.Print("Enter database password: ")
	fmt.Scanln(&inputPassword)

	if !authenticate(inputPassword) {
		fmt.Println("Incorrect password. Access denied.")
		return
	}

	if err := LoadEnv(".env"); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Get certificate and key paths
	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")

	if certFile == "" || keyFile == "" {
		log.Fatal("CERT_FILE and KEY_FILE environment variables must be set")
	}
	models.DatabaseExecution()
	defer models.CloseDatabase()
	routes.WebRouter()
	routes.ApiRouter()
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to load TLS certificate: %v", err)
	}
	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
		TLSConfig: &tls.Config{
			Certificates:             []tls.Certificate{cert}, // Explicitly set certificate
			MinVersion:               tls.VersionTLS12,
			PreferServerCipherSuites: true,
			CurvePreferences:         []tls.CurveID{tls.CurveP256, tls.X25519},
		},
	}

	fmt.Println("Server is running on https://localhost:8080")

	log.Fatal(server.ListenAndServeTLS("", ""))


}
