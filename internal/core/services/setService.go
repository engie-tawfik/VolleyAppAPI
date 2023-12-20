package services

import (
	"fmt"
	"time"
	"volleyapp/internal/core/constants"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
	"volleyapp/logger"
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

func (s *SetService) CreateSet(newSet domain.SetMainInfo) (int, error) {
	loc, _ := time.LoadLocation("America/Bogota")
	newSet.StartedAt = time.Now().In(loc)
	newSet.LastUpdate = time.Now().In(loc)
	newSet.IsActive = true
	setId, err := s.setRepository.SaveNewSet(newSet)
	if err != nil {
		errorMsg := fmt.Sprintf(
			"[SET SERVICE] Error in create set: %s", err,
		)
		return 0, fmt.Errorf(errorMsg)
	}
	return setId, nil
}

func (s *SetService) FinishSet(setId int) (int, error) {
	fail := func(err error) (int, error) {
		return 0, fmt.Errorf("[SET SERVICE] Error in finish set: %s", err)
	}
	set, err := s.setRepository.GetSet(setId)
	if err != nil {
		return fail(err)
	}
	gameTeamNames, err := s.gameRepository.GetTeamsNames(set.GameId)
	if err != nil {
		return fail(err)
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
		return fail(err)
	}
	return rowsAffected, nil
}

func (s *SetService) PlaySet(rally domain.Rally) (int, error) {
	forward := true
	fail := func(err error) (int, error) {
		return 0, fmt.Errorf("[SET SERVICE] Error in play set: %s", err)
	}
	set, err := s.setRepository.GetSet(rally.SetId)
	if err != nil {
		return fail(err)
	}
	logger.Logger.Debug(fmt.Sprintf("[SET SERVICE] Set from db: %v", set))
	if !set.IsActive {
		return fail(fmt.Errorf("set is not active"))
	}
	if rally.Action == constants.RollBack {
		if len(set.GameActions) > 0 {
			forward = false
			rally.Action = set.GameActions[len(set.GameActions)-1]
			set.GameActions = set.GameActions[:len(set.GameActions)-1]
		} else {
			return fail(fmt.Errorf("no registered actions in set"))
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
	logger.Logger.Debug(fmt.Sprintf("[SET SERVICE] Set to be saved: %v", set))
	_, err = s.setRepository.SaveSet(set)
	if err != nil {
		return fail(err)
	}
	rowsAffected, err := s.gameService.UpdateGameStats(set)
	if err != nil {
		return fail(err)
	}
	return rowsAffected, nil
}
