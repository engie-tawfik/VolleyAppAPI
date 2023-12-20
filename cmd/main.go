package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"volleyapp/internal/controllers"
	"volleyapp/internal/core/services"
	"volleyapp/internal/infrastructure/config"
	"volleyapp/internal/infrastructure/repositories"
	"volleyapp/internal/middlewares"
	"volleyapp/internal/server"
	"volleyapp/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.InitLogger()
	defer logger.StopLogger()

	if err := godotenv.Load(); err != nil {
		logger.Logger.Error(".env file not found")
	}

	instance := gin.New()
	instance.Use(gin.Recovery())

	dbDriver := os.Getenv("DB_DRIVER")
	dbUrl := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := repositories.NewDB(
		config.DatabaseConfig{
			Driver:                   dbDriver,
			Url:                      dbUrl,
			ConnMaxLifetimeInMinutes: 3,
			MaxOppenConns:            10,
			MaxIdleConns:             1,
		},
	)
	if err != nil {
		logger.Logger.Error(
			fmt.Sprintf(
				"Failed to connect to Database. Error: %s",
				err,
			),
		)
	}

	authMiddleware := middlewares.NewAuthMiddleware()
	headersMiddleware := middlewares.NewHeadersMiddleware()

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(
		instance,
		authService,
		authMiddleware,
		headersMiddleware,
	)
	authController.InitAuthRoutes()

	teamRepository := repositories.NewTeamRepository(db)
	teamService := services.NewTeamService(teamRepository)
	teamController := controllers.NewTeamController(
		instance,
		teamService,
		authMiddleware,
		headersMiddleware,
	)
	teamController.InitTeamRoutes()

	gameRepository := repositories.NewGameRepository(db)
	gameService := services.NewGameService(gameRepository)
	gameController := controllers.NewGameController(
		instance,
		gameService,
		authMiddleware,
		headersMiddleware,
	)
	gameController.InitGameRoutes()

	setRepository := repositories.NewSetRepository(db)
	setService := services.NewSetService(
		setRepository,
		gameRepository,
		gameService,
	)
	setController := controllers.NewSetController(
		instance,
		setService,
		authMiddleware,
		headersMiddleware,
	)
	setController.InitSetRoutes()

	httpServer := server.NewHttpServer(
		instance,
		config.HttpServerConfig{
			Port: 8080,
		},
	)

	httpServer.Start()
	defer httpServer.Stop()

	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-c
	logger.Logger.Info("See you in next game!")
}
