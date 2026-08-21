package main

import (
	"log"
	"os"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Backend server starting on port %s\n", port)
	// TODO: Inicializar Gin router e database connection
}
