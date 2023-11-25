package services

import (
	"fmt"
	"time"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
	"volleyapp/logger"
)

type TeamService struct {
	teamRepository ports.TeamRepository
}

var _ ports.TeamService = (*TeamService)(nil)

func NewTeamService(repository ports.TeamRepository) *TeamService {
	return &TeamService{
		teamRepository: repository,
	}
}

func (t *TeamService) CreateTeam(newTeam domain.TeamMainInfo) (int, error) {
	loc, _ := time.LoadLocation("America/Bogota")
	newTeam.UserId = 1
	newTeam.CreationDate = time.Now().In(loc)
	newTeam.LastUpdateDate = time.Now().In(loc)
	teamId, err := t.teamRepository.SaveNewTeam(newTeam)
	if err != nil {
		errorMsg := fmt.Sprintf(
			"[TEAM SERVICE] Error in create team: %s", err,
		)
		return 0, fmt.Errorf(errorMsg)
	}
	return teamId, nil
}

func (t *TeamService) GetUserTeams(userId int) ([]domain.TeamSummary, error) {
	userTeams, err := t.teamRepository.GetUserTeams(userId)
	if err != nil {
		errorMsg := fmt.Sprintf(
			"[TEAM SERVICE] Error in get user teams: %s", err,
		)
		return []domain.TeamSummary{}, fmt.Errorf(errorMsg)
	}
	return userTeams, nil
}

func (t *TeamService) GetTeam(teamId string) (domain.Team, error) {
	team, err := t.teamRepository.GetTeam(teamId)
	if err != nil {
		errorMsg := fmt.Sprintf(
			"TEAM SERVICE error GetTeam/GetTeam: %v",
			err,
		)
		logger.Logger.Error(errorMsg)
		return team, fmt.Errorf(errorMsg)
	}
	// These lines because mongo stored date in utc
	// loc, _ := time.LoadLocation("America/Bogota")
	// team.CreationDate = team.CreationDate.In(loc)
	return team, nil
}

func (t *TeamService) UpdateTeamInfo(team domain.TeamMainInfo) (bool, error) {
	return true, nil
}
