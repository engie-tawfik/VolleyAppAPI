package services

import (
	"fmt"
	"time"
	"volleyapp/domain/models"
	"volleyapp/domain/ports"
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

func (t *TeamService) CreateTeam(newTeam models.TeamMainInfo) (int, error) {
	loc, _ := time.LoadLocation("America/Bogota")
	newTeam.UserId = 1
	newTeam.CreationDate = time.Now().In(loc)
	newTeam.LastUpdateDate = time.Now().In(loc)
	teamId, err := t.teamRepository.SaveNewTeam(newTeam)
	if err != nil {
		return 0, fmt.Errorf("team service error in CreateTeam: %v", err)
	}
	return teamId, nil
}

func (t *TeamService) GetUserTeams(userId int) ([]models.TeamSummary, error) {
	userTeams, err := t.teamRepository.GetUserTeams(userId)
	if err != nil {
		return []models.TeamSummary{}, fmt.Errorf("team service error in GetUserTeams: %v", err)
	}
	return userTeams, nil
}

func (t *TeamService) GetTeam(teamId string) (models.Team, error) {
	team, err := t.teamRepository.GetTeam(teamId)
	if err != nil {
		return team, fmt.Errorf("team service error in GetTeam: %v", err)
	}
	return team, nil
}

func (t *TeamService) UpdateTeamInfo(team models.TeamMainInfo) (bool, error) {
	return true, nil
}
