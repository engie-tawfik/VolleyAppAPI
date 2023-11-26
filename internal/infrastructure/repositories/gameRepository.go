package repositories

import (
	"fmt"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
)

type GameRepository struct {
	db ports.Database
}

var _ ports.GameRepository = (*GameRepository)(nil)

func NewGameRepository(db ports.Database) *GameRepository {
	return &GameRepository{
		db: db,
	}
}

func (g *GameRepository) SaveNewGame(newGame domain.NewGame) (int, error) {
	query := "INSERT INTO game (team_id, game_date, is_active, game_country, game_province, game_city, opponent, last_update_date) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING game_id"
	result := g.db.GetDB().QueryRow(
		query,
		newGame.TeamId,
		newGame.GameDate,
		newGame.IsActive,
		newGame.GameCountry,
		newGame.GameProvince,
		newGame.GameCity,
		newGame.Opponent,
		newGame.LastUpdateDate,
	)
	var newGameId int
	if err := result.Scan(&newGameId); err != nil {
		return 0, fmt.Errorf(
			"[DATABASE] Error in save new game: %s", err,
		)
	}
	return int(newGameId), nil
}
