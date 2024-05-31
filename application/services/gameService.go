package services

import (
	"fmt"
	"log"
	"math"
	"time"
	"volleyapp/domain/models"
	"volleyapp/domain/ports"
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

func (g *GameService) CreateGame(newGame models.GameMainInfo) (int, error) {
	loc, _ := time.LoadLocation("America/Bogota")
	newGame.GameDate = time.Now().In(loc)
	newGame.LastUpdateDate = time.Now().In(loc)
	newGame.IsActive = true
	gameId, err := g.gameRepository.SaveNewGame(newGame)
	if err != nil {
		return 0, fmt.Errorf("game service - error in CreateGame: %v", err)
	}
	return gameId, nil
}

func (g *GameService) FinishGame(gameId int) (int, error) {
	loc, _ := time.LoadLocation("America/Bogota")
	game := models.GameMainInfo{
		LastUpdateDate: time.Now().In(loc),
		IsActive:       false,
	}
	rowsAffected, err := g.gameRepository.FinishGame(gameId, game)
	if err != nil {
		return 0, fmt.Errorf("game service - error in FinishGame: %v", err)
	}
	return rowsAffected, nil
}

func (g *GameService) UpdateGameStats(set models.Set) (int, error) {
	fail := func(err error) (int, error) {
		return 0, fmt.Errorf(
			"game service - error in UpdateGameStats: %v",
			err,
		)
	}
	game, err := g.gameRepository.GetGame(set.GameId)
	if err != nil {
		return fail(err)
	}
	log.Println("Game service - game from db:", game)
	if !game.IsActive {
		return fail(fmt.Errorf("game is not active"))
	}
	game.TotalAttacks += set.TotalAttacks
	game.AttackPoints += set.AttackPoints
	game.AttackNeutrals += set.AttackNeutrals
	game.TotalBlocks += set.TotalBlocks
	game.BlockPoints += set.BlockPoints
	game.BlockNeutrals += set.BlockNeutrals
	game.BlockErrors += set.BlockErrors
	game.TotalServes += set.TotalServes
	game.ServePoints += set.ServePoints
	game.ServeNeutrals += set.ServeNeutrals
	game.ServeErrors += set.ServeErrors
	game.OpponentErrors += set.OpponentErrors
	game.TotalPoints += set.TotalPoints
	game.OpponentAttacks += set.OpponentAttacks
	game.OpponentBlocks += set.OpponentBlocks
	game.OpponentServes += set.OpponentServes
	game.TotalErrors += set.TotalErrors
	game.OpponentPoints += set.OpponentPoints
	game.TotalActions += set.TotalActions
	game.AttackEffectiveness =
		(float64(game.AttackPoints) / float64(game.TotalAttacks)) * 100
	if math.IsNaN(game.AttackEffectiveness) {
		game.AttackEffectiveness = 0.00
	}
	game.TotalBlocks = game.BlockPoints + game.BlockNeutrals + game.BlockErrors
	game.BlockEffectiveness =
		(float64(game.BlockPoints) / float64(game.TotalBlocks)) * 100
	if math.IsNaN(game.BlockEffectiveness) {
		game.BlockEffectiveness = 0.00
	}
	game.TotalServes = game.ServePoints + game.ServeNeutrals + game.ServeErrors
	game.ServeEffectiveness =
		(float64(game.ServePoints) / float64(game.TotalServes)) * 100
	if math.IsNaN(game.ServeEffectiveness) {
		game.ServeEffectiveness = 0.00
	}
	game.TotalEffectiveness =
		(float64(game.TotalPoints-game.OpponentErrors) / float64(game.TotalActions)) * 100
	if math.IsNaN(game.TotalEffectiveness) {
		game.TotalEffectiveness = 0.00
	}

	loc, _ := time.LoadLocation("America/Bogota")
	game.LastUpdateDate = time.Now().In(loc)
	log.Println("Game service - game to be saved", game)
	rowsAffected, err := g.gameRepository.SaveGame(game)
	if err != nil {
		return fail(err)
	}
	// TODO update team stats
	return rowsAffected, nil
}
