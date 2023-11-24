package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"volleyapp/internal/infrastructure/config"

	"github.com/gin-gonic/gin"
)

const defaulHost = "localhost"

type HttpServer interface {
	Start()
	Stop()
}

type httpServer struct {
	// Add here every new Controller
	// authController ports.AuthController
	// teamController ports.TeamController
	Port   int
	server *http.Server
}

var _ HttpServer = (*httpServer)(nil)

func NewHttpServer(
	// authController ports.AuthController,
	// teamController ports.TeamController,
	router *gin.Engine,
	config config.HttpServerConfig,
) HttpServer {
	return &httpServer{
		Port: config.Port,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", defaulHost, config.Port),
			Handler: router,
		},
	}
}

func (h *httpServer) Start() {
	go func() {
		if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(
				"Failed to start HttpServer in port %d. Error: %s",
				h.Port,
				err.Error(),
			)
		}
	}()
	log.Printf("Server running in port %d", h.Port)
}

func (h *httpServer) Stop() {
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(3)*time.Second,
	)
	defer cancel()

	if err := h.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown. Error: %s", err.Error())
	}
}

// func (s *Server) Initialize() {
// 	app := gin.Default()
// 	basePath := app.Group("/v1/volleyapp")
// 	s.authController.RegisterAuthRoutes(basePath)
// 	s.teamController.RegisterTeamRoutes(basePath)
// 	app.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
// }
