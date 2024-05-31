package repositories

import (
	"fmt"
	"volleyapp/config"
	"volleyapp/domain/models"
	"volleyapp/domain/ports"
)

type GameRepository struct{}

var _ ports.GameRepository = (*GameRepository)(nil)

func NewGameRepository() *GameRepository {
	return &GameRepository{}
}

func (g *GameRepository) SaveNewGame(
	newGame models.GameMainInfo,
) (int, error) {
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
	result := config.DB.QueryRow(
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
			"database - error in SaveNewGame: %s", err,
		)
	}
	return int(newGameId), nil
}

func (g *GameRepository) FinishGame(
	gameId int,
	game models.GameMainInfo,
) (int, error) {
	query := `
		UPDATE game
		SET
			is_active = $1,
			last_update_date = $2
		WHERE game_id = $3
	`
	result, err := config.DB.Exec(
		query,
		game.IsActive,
		game.LastUpdateDate,
		gameId,
	)
	if err != nil {
		return 0, fmt.Errorf(
			"database - error in FinishGame: %s", err,
		)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf(
			"database - error in FinishGame: %s", err,
		)
	}
	return int(rowsAffected), nil
}

func (g *GameRepository) GetTeamsNames(
	gameId int,
) (models.GameTeamsNames, error) {
	var names models.GameTeamsNames
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
	result := config.DB.QueryRow(
		query,
		gameId,
	)
	if err := result.Scan(
		&names.OpponentName,
		&teamId,
		&names.TeamName,
	); err != nil {
		return names, fmt.Errorf(
			"database - error in GetTeamsNames: %v", err,
		)
	}
	return names, nil
}

func (g *GameRepository) GetGame(gameId int) (models.Game, error) {
	var game models.Game
	query := `
		SELECT *
		FROM game
		WHERE game_id = $1
	`
	result := config.DB.QueryRow(query, gameId)
	if err := result.Scan(
		&game.GameId,
		&game.TeamId,
		&game.GameDate,
		&game.IsActive,
		&game.GameCountry,
		&game.GameProvince,
		&game.GameCity,
		&game.Opponent,
		&game.TeamSets,
		&game.OpponentSets,
		&game.TotalAttacks,
		&game.AttackPoints,
		&game.AttackNeutrals,
		&game.AttackErrors,
		&game.AttackEffectiveness,
		&game.TotalBlocks,
		&game.BlockPoints,
		&game.BlockNeutrals,
		&game.BlockErrors,
		&game.BlockEffectiveness,
		&game.TotalServes,
		&game.ServePoints,
		&game.ServeNeutrals,
		&game.ServeErrors,
		&game.ServeEffectiveness,
		&game.OpponentErrors,
		&game.TotalPoints,
		&game.TotalActions,
		&game.TotalEffectiveness,
		&game.OpponentAttacks,
		&game.OpponentBlocks,
		&game.OpponentServes,
		&game.TotalErrors,
		&game.OpponentPoints,
		&game.GameWinner,
		&game.LastUpdateDate,
	); err != nil {
		return game, fmt.Errorf("database - error in GetGame: %v", err)
	}
	return game, nil
}

func (g *GameRepository) SaveGame(game models.Game) (int, error) {
	query := `
		UPDATE game
		SET
			team_sets = $1,
			opponent_sets = $2,
			total_attacks = $3,
			attack_points = $4,
			attack_neutrals = $5,
			attack_errors = $6,
			attack_effectiveness = $7,
			total_blocks = $8,
			block_points = $9,
			block_neutrals = $10,
			block_errors = $11,
			block_effectiveness = $12,
			total_serves = $13,
			serve_points = $14,
			serve_neutrals = $15,
			serve_errors = $16,
			serve_effectiveness = $17,
			opponent_errors = $18,
			total_points = $19,
			total_actions = $20,
			total_effectiveness = $21,
			opponent_attacks = $22,
			opponent_blocks = $23,
			opponent_serves = $24,
			total_errors = $25,
			opponent_points = $26,
			last_update_date = $27
		WHERE game_id = $28
	`
	result, err := config.DB.Exec(
		query,
		game.TeamSets,
		game.OpponentSets,
		game.TotalAttacks,
		game.AttackPoints,
		game.AttackNeutrals,
		game.AttackErrors,
		game.AttackEffectiveness,
		game.TotalBlocks,
		game.BlockPoints,
		game.BlockNeutrals,
		game.BlockErrors,
		game.BlockEffectiveness,
		game.TotalServes,
		game.ServePoints,
		game.ServeNeutrals,
		game.ServeErrors,
		game.ServeEffectiveness,
		game.OpponentErrors,
		game.TotalPoints,
		game.TotalActions,
		game.TotalEffectiveness,
		game.OpponentAttacks,
		game.OpponentBlocks,
		game.OpponentServes,
		game.TotalErrors,
		game.OpponentPoints,
		game.LastUpdateDate,
		game.GameId,
	)
	if err != nil {
		return 0, fmt.Errorf(
			"database - error in SaveGame: %v", err,
		)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf(
			"database - error in SaveGame: %v", err,
		)
	}
	return int(rowsAffected), nil
}
