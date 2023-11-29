package repositories

import (
	"fmt"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
)

type SetRepository struct {
	db ports.Database
}

var _ ports.SetRepository = (*SetRepository)(nil)

func NewSetRepository(db ports.Database) *SetRepository {
	return &SetRepository{
		db: db,
	}
}

func (s *SetRepository) SaveNewSet(newSet domain.SetMainInfo) (int, error) {
	query := "INSERT INTO set (game_id, started_at, is_active, last_update) VALUES($1, $2, $3, $4) RETURNING set_id"
	result := s.db.GetDB().QueryRow(
		query,
		newSet.GameId,
		newSet.StartedAt,
		newSet.IsActive,
		newSet.LastUpdate,
	)
	var newSetId int
	if err := result.Scan(&newSetId); err != nil {
		return 0, fmt.Errorf(
			"[DATABASE] Error in save new set: %s", err,
		)
	}
	return int(newSetId), nil
}

func (s *SetRepository) FinishSet(setId int, set domain.SetMainInfo) (int, error) {
	query := "UPDATE set SET is_active = $1, last_update = $2 WHERE set_id = $3"
	result, err := s.db.GetDB().Exec(
		query,
		set.IsActive,
		set.LastUpdate,
		setId,
	)
	if err != nil {
		return 0, fmt.Errorf(
			"[DATABASE] Error in finish set: %s", err,
		)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf(
			"[DATABASE] Error in finish set: %s", err,
		)
	}
	return int(rowsAffected), nil
}
