package main

import (
	"crypto/tls"
	"fmt"
	"log"

	"forum/app/models"
	"forum/routes"
	"net/http"

)


func main() {
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
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))


	fmt.Println("Server is running on https://localhost:8443")
	
}
