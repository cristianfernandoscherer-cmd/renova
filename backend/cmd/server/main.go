package main

import (
	"log"
	"os"
)

// Dependências instaladas em go.mod (bootstrap inicial):
// - github.com/gin-gonic/gin (HTTP router)
// - gorm.io/gorm + gorm.io/driver/postgres (ORM e PostgreSQL)
// - github.com/golang-jwt/jwt/v5 (JWT tokens)
// - golang.org/x/crypto (bcrypt para senhas)
// - github.com/cosmtrek/air (hot reload dev)
// - github.com/stretchr/testify (testing)
//
// Estas dependências não são importadas ainda pois TASK-005 é apenas
// o bootstrap do projeto. Elas serão utilizadas em tasks subsequentes.

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Backend server starting on port %s\n", port)
	// TODO: Inicializar Gin router e database connection
}
