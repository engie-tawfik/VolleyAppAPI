package main

import (
	"context"
	"volleyapp/config"
	"volleyapp/internal/core/services"
	"volleyapp/internal/database"
	"volleyapp/internal/handlers"
	"volleyapp/internal/middlewares"
	"volleyapp/internal/repositories"
	"volleyapp/internal/server"
)

func init() {
	config.LoadEnvs()
	config.InitLogger()
	database.Connect()
}

func main() {
	ctx := context.TODO()

	// Repositories
	authRepository := repositories.NewAuthRepository(database.Collection, ctx)
	teamRepository := repositories.NewTeamRepository(database.Collection, ctx)
	// Services
	authService := services.NewAuthService(authRepository)
	teamService := services.NewTeamService(teamRepository)
	// Middlewares
	authMiddleware := middlewares.NewAuthMiddleware()
	// Handlers
	authHandler := handlers.NewAuthHandler(authService, authMiddleware)
	teamHandler := handlers.NewTeamHandler(teamService, authMiddleware)
	// Server
	ginServer := server.NewServer(authHandler, teamHandler)
	ginServer.Initialize()

	defer database.Disconnect(ctx)
	defer config.StopLogger()
}
