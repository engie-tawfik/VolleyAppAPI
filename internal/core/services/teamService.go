package services

import (
	"time"
	"volleyapp/config"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
	"volleyapp/utils"
)

type TeamService struct {
	teamRepository ports.TeamRepository
}

var _ ports.TeamService = (*TeamService)(nil)

func NewTeamService(teamRepository ports.TeamRepository) *TeamService {
	return &TeamService{
		teamRepository: teamRepository,
	}
}

func (t *TeamService) CreateTeam(team domain.NewTeam) string {
	teamExists, err := t.teamRepository.CheckTeamExistence(team.Email)
	if err != nil {
		return "No connection with database."
	}
	if teamExists {
		config.Logger.Info("Team already registered in database")
		return "Team already registered in database."
	}
	hashedPass := utils.Hash(team.Password)
	if hashedPass == "" {
		return "Unable to create team."
	}
	team.Password = hashedPass
	loc, _ := time.LoadLocation("America/Bogota")
	team.CreationDateTime = time.Now().In(loc)
	result, ok := t.teamRepository.CreateTeam(team)
	if !ok {
		return result
	}
	config.Logger.Info("New Team created with id " + result)
	return ""
}

func (t *TeamService) GetTeam(teamId string) domain.Team {
	team := t.teamRepository.GetTeam(teamId)
	if team.Name == "" {
		return team
	}
	loc, _ := time.LoadLocation("America/Bogota")
	team.CreationDateTime = team.CreationDateTime.In(loc)
	return team
}
