package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go_yoga_api/internal/handler"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	// creates and returns a new Gin engine instance
	// which is the main router responsible for handling HTTP requests and routing them to the appropriate handlers

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)

	r.GET("/api/v1/poses", handler.GetAllPosesHandler)
	r.GET("/api/v1/poses/:id", handler.GetOnePoseHandler)

	r.GET("/api/v1/routines", handler.GetAllRoutinesHandler)
	r.GET("/api/v1/routines/:id", handler.GetOneRoutineHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
