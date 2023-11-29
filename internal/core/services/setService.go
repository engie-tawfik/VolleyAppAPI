package services

import (
	"fmt"
	"time"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
)

type SetService struct {
	setRepository ports.SetRepository
}

var _ ports.SetService = (*SetService)(nil)

func NewSetService(repository ports.SetRepository) *SetService {
	return &SetService{
		setRepository: repository,
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
	loc, _ := time.LoadLocation("America/Bogota")
	set := domain.SetMainInfo{
		LastUpdate: time.Now().In(loc),
		IsActive:   false,
	}
	rowsAffected, err := s.setRepository.FinishSet(setId, set)
	if err != nil {
		errorMsg := fmt.Sprintf(
			"[SET SERVICE] Error in finish set: %s", err,
		)
		return 0, fmt.Errorf(errorMsg)
	}
	return rowsAffected, nil
}
