package repositories

import (
	"fmt"
	"volleyapp/config"
	"volleyapp/domain/models"
	"volleyapp/domain/ports"
)

type TeamRepository struct{}

var _ ports.TeamRepository = (*TeamRepository)(nil)

func NewTeamRepository() *TeamRepository {
	return &TeamRepository{}
}

func (t *TeamRepository) CheckTeamExistence(email string) (bool, error) {
	return true, nil
}

func (t *TeamRepository) SaveNewTeam(newTeam models.TeamMainInfo) (int, error) {
	query := `
		INSERT INTO teams
			(user_id, team_name, team_country, team_province, team_city, team_category, creation_date, last_update_date)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING team_id
	`
	result := config.DB.QueryRow(
		query,
		newTeam.UserId,
		newTeam.Name,
		newTeam.Country,
		newTeam.Province,
		newTeam.City,
		newTeam.Category,
		newTeam.CreationDate,
		newTeam.LastUpdateDate,
	)
	var newTeamId int
	if err := result.Scan(&newTeamId); err != nil {
		return 0, fmt.Errorf(
			"database - error in SaveNewTeam: %v", err,
		)
	}
	return int(newTeamId), nil
}

func (t *TeamRepository) GetUserTeams(
	userId int,
) ([]models.TeamSummary, error) {
	var teams []models.TeamSummary
	query := `
		SELECT
			team_id,
			user_id,
			team_name,
			team_country,
			team_province,
			team_city,
			team_category,
			creation_date,
			last_update_date,
			total_games,
			won_games,
			total_sets,
			won_sets,
			attack_effectiveness,
			block_effectiveness,
			serve_effectiveness,
			total_effectiveness
		FROM teams
		WHERE user_id = $1
	`
	result, err := config.DB.Query(query, userId)
	if err != nil {
		return []models.TeamSummary{}, fmt.Errorf(
			"database - error in GetUserTeams: %v",
			err,
		)
	}
	defer result.Close()
	for result.Next() {
		var team models.TeamSummary
		if err := result.Scan(
			&team.TeamMainInfo.TeamId,
			&team.TeamMainInfo.UserId,
			&team.TeamMainInfo.Name,
			&team.TeamMainInfo.Country,
			&team.TeamMainInfo.Province,
			&team.TeamMainInfo.City,
			&team.TeamMainInfo.Category,
			&team.TeamMainInfo.CreationDate,
			&team.TeamMainInfo.LastUpdateDate,
			&team.TotalGames,
			&team.WonGames,
			&team.TotalSets,
			&team.WonSets,
			&team.AttackEffectiveness,
			&team.BlockEffectiveness,
			&team.ServeEffectiveness,
			&team.TotalEffectiveness,
		); err != nil {
			return []models.TeamSummary{}, fmt.Errorf(
				"database - error in GetUSerTeams: %v",
				err,
			)
		}
		teams = append(teams, team)
	}
	if err := result.Err(); err != nil {
		return []models.TeamSummary{}, fmt.Errorf(
			"database - error in get user teams: %v",
			err,
		)
	}
	return teams, nil
}

func (t *TeamRepository) GetTeam(teamId int) (models.Team, error) {
	var team models.Team
	query := `
		SELECT
			team_id,
			user_id,
			team_name,
			team_country,
			team_province,
			team_city,
			team_category,
			total_games,
			won_games,
			total_sets,
			won_sets,
			total_attacks,
			attack_points,
			attack_neutrals,
			attack_errors,
			attack_effectiveness,
			total_blocks,
			block_points,
			block_neutrals,
			block_errors,
			block_effectiveness,
			total_serves,
			serve_points,
			serve_neutrals,
			serve_errors,
			serve_effectiveness,
			opponent_errors,
			total_points,
			total_actions,
			total_effectiveness,
			opponent_attacks,
			opponent_blocks,
			opponent_serves,
			total_errors,
			creation_date,
			last_update_date
		FROM teams
		WHERE team_id = $1
		`
	result := config.DB.QueryRow(
		query,
		teamId,
	)
	if err := result.Scan(
		&team.TeamMainInfo.TeamId,
		&team.TeamMainInfo.UserId,
		&team.TeamMainInfo.Name,
		&team.TeamMainInfo.Country,
		&team.TeamMainInfo.Province,
		&team.TeamMainInfo.City,
		&team.TeamMainInfo.Category,
		&team.TeamGameData.TotalGames,
		&team.TeamGameData.WonGames,
		&team.TeamGameData.TotalSets,
		&team.TeamGameData.WonSets,
		&team.TeamGameData.TotalAttacks,
		&team.TeamGameData.AttackPoints,
		&team.TeamGameData.AttackNeutrals,
		&team.TeamGameData.AttackErrors,
		&team.TeamGameData.AttackEffectiveness,
		&team.TeamGameData.TotalBlocks,
		&team.TeamGameData.BlockPoints,
		&team.TeamGameData.BlockNeutrals,
		&team.TeamGameData.BlockErrors,
		&team.TeamGameData.BlockEffectiveness,
		&team.TeamGameData.TotalServes,
		&team.TeamGameData.ServePoints,
		&team.TeamGameData.ServeNeutrals,
		&team.TeamGameData.ServeErrors,
		&team.TeamGameData.ServeEffectiveness,
		&team.TeamGameData.OpponentErrors,
		&team.TeamGameData.TotalPoints,
		&team.TeamGameData.TotalActions,
		&team.TeamGameData.TotalEffectiveness,
		&team.TeamGameData.OpponentAttacks,
		&team.TeamGameData.OpponentBlocks,
		&team.TeamGameData.OpponentServes,
		&team.TeamGameData.TotalErrors,
		&team.TeamMainInfo.CreationDate,
		&team.TeamMainInfo.LastUpdateDate,
	); err != nil {
		return team, fmt.Errorf(
			"database - error in GetTeam: %v", err,
		)
	}
	return team, nil
}

func (t *TeamRepository) UpdateTeamInfo(
	team models.TeamMainInfo,
) (bool, error) {
	return true, nil
}
