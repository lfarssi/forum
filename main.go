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

func main() {

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
	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS12, // Enforce TLS 1.2+
			PreferServerCipherSuites: true,
			CurvePreferences:         []tls.CurveID{tls.CurveP256, tls.X25519},
		},
	}

	log.Println("Starting HTTPS server on port 8080...")
	log.Fatal(server.ListenAndServeTLS(certFile, keyFile))

	fmt.Println("Server is running on https://localhost:8443")

}
