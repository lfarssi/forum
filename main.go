package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"forum/app/models"
	"forum/routes"
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
			continue
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

	tlsConfig := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
		PreferServerCipherSuites: true,
	}

	server := &http.Server{
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
		TLSConfig:    tlsConfig,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Println("Server is running on https://localhost:8080")
	server.ListenAndServeTLS(certFile, keyFile)
}
