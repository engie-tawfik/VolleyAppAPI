package routes

import (
	"volleyapp/application/services"
	"volleyapp/config"
	"volleyapp/infrastructure/controllers"
	"volleyapp/infrastructure/middlewares"
	"volleyapp/infrastructure/repositories"
)

func InitRoutes() {
	authMiddleware := middlewares.NewAuthMiddleware()
	headersMiddleware := middlewares.NewHeadersMiddleware()

	authRepository := repositories.NewAuthRepository()
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(
		config.Server,
		authService,
		authMiddleware,
		headersMiddleware,
	)
	authController.InitAuthRoutes()

	teamRepository := repositories.NewTeamRepository()
	teamService := services.NewTeamService(teamRepository)
	teamController := controllers.NewTeamController(
		config.Server,
		teamService,
		authMiddleware,
		headersMiddleware,
	)
	teamController.InitTeamRoutes()

	gameRepository := repositories.NewGameRepository()
	gameService := services.NewGameService(gameRepository)
	gameController := controllers.NewGameController(
		config.Server,
		gameService,
		authMiddleware,
		headersMiddleware,
	)
	gameController.InitGameRoutes()

	setRepository := repositories.NewSetRepository()
	setService := services.NewSetService(
		setRepository,
		gameRepository,
		gameService,
	)
	setController := controllers.NewSetController(
		config.Server,
		setService,
		authMiddleware,
		headersMiddleware,
	)
	setController.InitSetRoutes()
}
