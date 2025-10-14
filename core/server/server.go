package server

import (
	"api/core/config"
	"api/core/repository"
	"api/core/routes"
	"log"
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	db     repository.DBClient
	config *config.Config
	// s3 bucket if you need it
	// validation *validator.Validate
	// translate *ut.Translator
}

func (s *Server) Start() error {
	slog.Info("Server starting at port 8080")
	err := s.router.Run(":8080")
	if err != nil {
		log.Fatalf("Server Issue: %s", err)
		return err
	}
	return nil
}

func NewServer(db repository.DBClient, cfg *config.Config) *Server {
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	server := Server{
		router: gin.Default(),
		db:     db,
		config: cfg,
	}

	server.router.Use(cors.New(config))
	routes.RegisterRoutes(server.router, cfg)
	return &server
}
