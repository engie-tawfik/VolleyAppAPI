package services

import (
	"fmt"
	"log"
	"time"
	"volleyapp/domain/constants"
	"volleyapp/domain/models"
	"volleyapp/domain/ports"
)

type SetService struct {
	setRepository  ports.SetRepository
	gameRepository ports.GameRepository
	gameService    ports.GameService
}

var _ ports.SetService = (*SetService)(nil)

func NewSetService(
	repository ports.SetRepository,
	gameRepository ports.GameRepository,
	gameService ports.GameService,
) *SetService {
	return &SetService{
		setRepository:  repository,
		gameRepository: gameRepository,
		gameService:    gameService,
	}
}

func (s *SetService) CreateSet(newSet models.SetMainInfo) (int, error) {
	loc, _ := time.LoadLocation("America/Bogota")
	newSet.StartedAt = time.Now().In(loc)
	newSet.LastUpdate = time.Now().In(loc)
	newSet.IsActive = true
	setId, err := s.setRepository.SaveNewSet(newSet)
	if err != nil {
		return 0, fmt.Errorf("set service - error in CreateSet: %v", err)
	}
	return setId, nil
}

func (s *SetService) FinishSet(setId int) (int, error) {
	handleError := func(err error) (int, error) {
		return 0, fmt.Errorf("set service - error in FinishSet: %v", err)
	}
	set, err := s.setRepository.GetSet(setId)
	if err != nil {
		return handleError(err)
	}
	gameTeamNames, err := s.gameRepository.GetTeamsNames(set.GameId)
	if err != nil {
		return handleError(err)
	}
	if set.TotalPoints > set.OpponentPoints {
		set.SetWinner = gameTeamNames.TeamName
	} else if set.TotalPoints < set.OpponentPoints {
		set.SetWinner = gameTeamNames.OpponentName
	}
	loc, _ := time.LoadLocation("America/Bogota")
	set.LastUpdate = time.Now().In(loc)
	set.IsActive = false
	rowsAffected, err := s.setRepository.FinishSet(setId, set)
	if err != nil {
		return handleError(err)
	}
	return rowsAffected, nil
}

func (s *SetService) PlaySet(rally models.Rally) (int, error) {
	forward := true
	handleError := func(err error) (int, error) {
		return 0, fmt.Errorf("set service - error in PlaySet: %v", err)
	}
	set, err := s.setRepository.GetSet(rally.SetId)
	if err != nil {
		return handleError(err)
	}
	log.Println("Set service - set from db:", set)
	if !set.IsActive {
		return handleError(fmt.Errorf("set is not active"))
	}
	if rally.Action == constants.RollBack {
		if len(set.GameActions) > 0 {
			forward = false
			rally.Action = set.GameActions[len(set.GameActions)-1]
			set.GameActions = set.GameActions[:len(set.GameActions)-1]
		} else {
			return handleError(fmt.Errorf("no registered actions in set"))
		}
	} else {
		set.GameActions = append(set.GameActions, rally.Action)
	}

	switch rally.Action {
	case constants.AttackPoint:
		set.AttackPoint(forward)
	case constants.AttackNeutral:
		set.AttackNeutral(forward)
	case constants.AttackError:
		set.AttackError(forward)
	case constants.OpponentAttack:
		set.OpponentAttack(forward)
	case constants.BlockPoint:
		set.BlockPoint(forward)
	case constants.BlockNeutral:
		set.BlockNeutral(forward)
	case constants.BlockError:
		set.BlockError(forward)
	case constants.OpponentBlock:
		set.OpponentBlock(forward)
	case constants.ServePoint:
		set.ServePoint(forward)
	case constants.ServeNeutral:
		set.ServeNeutral(forward)
	case constants.ServeError:
		set.ServeError(forward)
	case constants.OpponentService:
		set.OpponentServe(forward)
	case constants.Error:
		set.Error(forward)
	case constants.OpponentError:
		set.OpponentError(forward)
	}
	set.UpdateStats()

	loc, _ := time.LoadLocation("America/Bogota")
	set.LastUpdate = time.Now().In(loc)
	log.Println("Set service - set to be saved:", set)
	_, err = s.setRepository.SaveSet(set)
	if err != nil {
		return handleError(err)
	}
	rowsAffected, err := s.gameService.UpdateGameStats(set)
	if err != nil {
		return handleError(err)
	}
	return rowsAffected, nil
}
