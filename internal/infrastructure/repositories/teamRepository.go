package repositories

import (
	"database/sql"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
)

type TeamRepository struct {
	db *sql.DB
}

var _ ports.TeamRepository = (*TeamRepository)(nil)

func NewTeamRepository(db *sql.DB) *TeamRepository {
	return &TeamRepository{
		db: db,
	}
}

func (t *TeamRepository) CheckTeamExistence(email string) (bool, error) {
	return true, nil
}

func (t *TeamRepository) CreateTeam(team domain.NewTeam) (bool, error) {
	return true, nil
}

func (t *TeamRepository) GetTeam(teamId string) (domain.Team, error) {
	var team domain.Team
	return team, nil
}

func (t *TeamRepository) UpdateTeamInfo(team domain.BaseTeam) (bool, error) {
	return true, nil
}
