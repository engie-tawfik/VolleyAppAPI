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
		errorMsg := fmt.Errorf(
			"TEAM SERVICE error CreateTeam/CheckTeamExistence: %v",
			err,
		)
		logger.Logger.Error(errorMsg.Error())
		return false, errorMsg
	}
	if teamExists {
		errorMsg := fmt.Errorf(
			"TEAM SERVICE error: team already registered in database",
		)
		logger.Logger.Info(errorMsg.Error())
		return false, errorMsg
	}
	hashedPass := utils.Hash(team.Password)
	if hashedPass == "" {
		errorMsg := fmt.Errorf(
			"TEAM SERVICE error: unable to hash password",
		)
		logger.Logger.Info(errorMsg.Error())
		return false, errorMsg
	}
	team.Password = hashedPass
	loc, _ := time.LoadLocation("America/Bogota")
	team.CreationDateTime = time.Now().In(loc)
	team.LastUpdateDateTime = time.Now().In(loc)
	_, err = t.teamRepository.CreateTeam(team)
	if err != nil {
		errorMsg := fmt.Errorf(
			"TEAM SERVICE error CreateTeam/CreateTeam: %v",
			err,
		)
		logger.Logger.Info(errorMsg.Error())
		return false, errorMsg
	}
	logger.Logger.Info("New Team successfully created")
	return true, nil
}

func (t *TeamService) GetTeam(teamId string) (domain.Team, error) {
	team, err := t.teamRepository.GetTeam(teamId)
	if err != nil {
		errorMsg := fmt.Errorf(
			"TEAM SERVICE error GetTeam/GetTeam: %v",
			err,
		)
		logger.Logger.Info(errorMsg.Error())
		return team, errorMsg
	}
	loc, _ := time.LoadLocation("America/Bogota")
	team.CreationDateTime = team.CreationDateTime.In(loc)
	return team, nil
}

func (t *TeamService) UpdateTeamInfo(team domain.BaseTeam) (bool, error) {
	return true, nil
}
