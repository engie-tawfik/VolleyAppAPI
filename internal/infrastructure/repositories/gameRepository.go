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

func (g *GameRepository) SaveNewGame(newGame domain.GameMainInfo) (int, error) {
	query := `
		INSERT INTO game (
			team_id,
			game_date,
			is_active,
			game_country,
			game_province,
			game_city,
			opponent,
			last_update_date
		) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING game_id`
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

func (g *GameRepository) FinishGame(
	gameId int,
	game domain.GameMainInfo,
) (int, error) {
	query := `
		UPDATE game
		SET 
			is_active = $1,
			last_update_date = $2
		WHERE game_id = $3
	`
	result, err := g.db.GetDB().Exec(
		query,
		game.IsActive,
		game.LastUpdateDate,
		gameId,
	)
	if err != nil {
		return 0, fmt.Errorf(
			"[DATABASE] Error in finish game: %s", err,
		)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf(
			"[DATABASE] Error in finish game: %s", err,
		)
	}
	return int(rowsAffected), nil
}

func (g *GameRepository) GetTeamsNames(
	gameId int,
) (domain.GameTeamsNames, error) {
	var names domain.GameTeamsNames
	var teamId int
	query := `
		SELECT
			g.opponent,
			g.team_id,
			t.team_name
		FROM game g
		JOIN team t
		ON t.team_id = g.team_id
		AND g.game_id = $1
	`
	result := g.db.GetDB().QueryRow(
		query,
		gameId,
	)
	if err := result.Scan(
		&names.OpponentName,
		&teamId,
		&names.TeamName,
	); err != nil {
		return names, fmt.Errorf(
			"[DATABASE] Error in get teams names: %s", err,
		)
	}
	return names, nil
}
