package services

import (
	"fmt"
	"time"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
	"volleyapp/logger"
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

func (t *TeamService) CreateTeam(team domain.NewTeam) (bool, error) {
	teamExists, err := t.teamRepository.CheckTeamExistence(team.Email)
	if err != nil {
		errorMsg := fmt.Sprintf(
			"TEAM SERVICE error CreateTeam/CheckTeamExistence: %s", err,
		)
		logger.Logger.Error(errorMsg)
		return false, fmt.Errorf(errorMsg)
	}
	if teamExists {
		errorMsg := "TEAM SERVICE error: team already registered in database"
		logger.Logger.Error(errorMsg)
		return false, fmt.Errorf(errorMsg)
	}
	hashedPass := utils.Hash(team.Password)
	if hashedPass == "" {
		errorMsg := "TEAM SERVICE error: unable to hash password"
		logger.Logger.Error(errorMsg)
		return false, fmt.Errorf(errorMsg)
	}
	team.Password = hashedPass
	loc, _ := time.LoadLocation("America/Bogota")
	team.CreationDateTime = time.Now().In(loc)
	team.LastUpdateDateTime = time.Now().In(loc)
	_, err = t.teamRepository.CreateTeam(team)
	if err != nil {
		errorMsg := fmt.Sprintf(
			"TEAM SERVICE error CreateTeam/CreateTeam: %s",
			err,
		)
		logger.Logger.Error(errorMsg)
		return false, fmt.Errorf(errorMsg)
	}
	logger.Logger.Info("New Team successfully created")
	return true, nil
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
	loc, _ := time.LoadLocation("America/Bogota")
	team.CreationDateTime = team.CreationDateTime.In(loc)
	return team, nil
}

func (t *TeamService) UpdateTeamInfo(team domain.BaseTeam) (bool, error) {
	return true, nil
}
