package repositories

import (
	"fmt"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
)

type TeamRepository struct {
	db ports.Database
}

var _ ports.TeamRepository = (*TeamRepository)(nil)

func NewTeamRepository(db ports.Database) *TeamRepository {
	return &TeamRepository{
		db: db,
	}
}

func (t *TeamRepository) CheckTeamExistence(email string) (bool, error) {
	return true, nil
}

func (t *TeamRepository) SaveNewTeam(newTeam domain.TeamMainInfo) (int, error) {
	query := "INSERT INTO team (user_id, team_name, team_country, team_province, team_city, team_category, creation_date, last_update_date) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING team_id"
	result := t.db.GetDB().QueryRow(
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
			"[DATABASE] Error in save new team: %s", err,
		)
	}
	return int(newTeamId), nil
}

func (t *TeamRepository) GetUserTeams(userId int) ([]domain.TeamSummary, error) {
	var teams []domain.TeamSummary
	query := "SELECT team_id, user_id, team_name, team_country, team_province, team_city, team_category, creation_date, last_update_date, total_games, won_games, total_sets, won_sets, attack_effectiveness, block_effectiveness, serve_effectiveness, total_effectiveness FROM team WHERE user_id = $1"
	result, err := t.db.GetDB().Query(query, userId)
	if err != nil {
		errorMsg := fmt.Sprintf("[DATABASE] Error in get user teams: %s", err)
		return []domain.TeamSummary{}, fmt.Errorf(errorMsg)
	}
	defer result.Close()
	for result.Next() {
		var team domain.TeamSummary
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
			errorMsg := fmt.Sprintf("[DATABASE] Error in get user teams: %s", err)
			return []domain.TeamSummary{}, fmt.Errorf(errorMsg)
		}
		teams = append(teams, team)
	}
	if err := result.Err(); err != nil {
		errorMsg := fmt.Sprintf("[DATABASE] Error in get user teams: %s", err)
		return []domain.TeamSummary{}, fmt.Errorf(errorMsg)
	}
	return teams, nil
}

func (t *TeamRepository) GetTeam(teamId string) (domain.Team, error) {
	var team domain.Team
	return team, nil
}

func (t *TeamRepository) UpdateTeamInfo(team domain.TeamMainInfo) (bool, error) {
	return true, nil
}
