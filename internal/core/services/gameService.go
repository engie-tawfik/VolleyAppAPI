package services

import (
	"fmt"
	"time"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
)

type GameService struct {
	gameRepository ports.GameRepository
}

var _ ports.GameService = (*GameService)(nil)

func NewGameService(repository ports.GameRepository) *GameService {
	return &GameService{
		gameRepository: repository,
	}
}

func (g *GameService) CreateGame(newGame domain.GameMainInfo) (int, error) {
	loc, _ := time.LoadLocation("America/Bogota")
	newGame.GameDate = time.Now().In(loc)
	newGame.LastUpdateDate = time.Now().In(loc)
	newGame.IsActive = true
	gameId, err := g.gameRepository.SaveNewGame(newGame)
	if err != nil {
		errorMsg := fmt.Sprintf(
			"[GAME SERVICE] Error in create game: %s", err,
		)
		return 0, fmt.Errorf(errorMsg)
	}
	return gameId, nil
}

func (g *GameService) FinishGame(gameId int) (int, error) {
	loc, _ := time.LoadLocation("America/Bogota")
	game := domain.GameMainInfo{
		LastUpdateDate: time.Now().In(loc),
		IsActive:       false,
	}
	rowsAffected, err := g.gameRepository.FinishGame(gameId, game)
	if err != nil {
		errorMsg := fmt.Sprintf(
			"[GAME SERVICE] Error in finish game: %s", err,
		)
		return 0, fmt.Errorf(errorMsg)
	}
	return rowsAffected, nil
}
