package server

import (
	"fmt"
	"os"
	"volleyapp/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type Server struct {
	// Add here every new Handler
	authHandler ports.AuthHandler
	teamHandler ports.TeamHandler
}

func NewServer(
	authHandler ports.AuthHandler,
	teamHandler ports.TeamHandler,
) *Server {
	return &Server{
		authHandler: authHandler,
		teamHandler: teamHandler,
	}
}

func (s *Server) Initialize() {
	app := gin.Default()
	basePath := app.Group("/v1/volleyapp")
	s.authHandler.RegisterAuthRoutes(basePath)
	s.teamHandler.RegisterTeamRoutes(basePath)
	app.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
