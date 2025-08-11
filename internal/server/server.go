package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"go_yoga_api/internal/database"
)

type Server struct {
	port int
	// ^ address

	db database.Service
	// ^ database connection
}
// NewServer creates a new HTTP server with specified configurations.
func NewServer() *http.Server {
	// Check enviroment variable PORT to set server port
	// Default port is 8080
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,

		db: database.New(),
		// ^ grabs a new database connection
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		// ^ RegisterRoutes is in routes.go
		// tells the server which routes it has
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
	// ^ starts the HTTP server
}
