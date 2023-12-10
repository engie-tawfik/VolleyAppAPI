package repositories

import (
	"context"
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

func (s *SetRepository) SaveRally(rally domain.Rally) (int, error) {
	fail := func(err error) (int, error) {
		return 0, fmt.Errorf("[DATABASE] Error in play set: %s", err)
	}
	ctx := context.TODO()
	tx, err := s.db.GetDB().BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}

	defer tx.Rollback()

	row := fmt.Sprintf("%ss", rally.Action)
	query := fmt.Sprintf(
		"UPDATE set SET %s = %s + 1, game_actions = ARRAY_APPEND(game_actions, $1), last_update = $2 WHERE set_id = $3",
		row,
		row,
	)
	result, err := tx.ExecContext(
		ctx,
		query,
		rally.Action,
		rally.DateTime,
		rally.SetId,
	)
	if err != nil {
		return fail(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil || int(rowsAffected) != 1 {
		return fail(err)
	}
	// No extra writings in set table for:
	// attack neutral
	// serve neutral

	// attack point - total points
	// block point - total points
	// serve point - total points
	// opponent error - total points
	if rally.Action == "attack_point" ||
		rally.Action == "block_point" ||
		rally.Action == "serve_point" ||
		rally.Action == "opponent_error" {
		query := "UPDATE set SET total_points = total_points + 1 WHERE set_id = $1"
		result, err := tx.ExecContext(
			ctx,
			query,
			rally.SetId,
		)
		if err != nil {
			return fail(err)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil || int(rowsAffected) != 1 {
			return fail(err)
		}
	}

	// block neutral - total blocks
	if rally.Action == "block_neutral" {
		query := "UPDATE set SET total_blocks = total_blocks + 1 WHERE set_id = $1"
		result, err := tx.ExecContext(
			ctx,
			query,
			rally.SetId,
		)
		if err != nil {
			return fail(err)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil || int(rowsAffected) != 1 {
			return fail(err)
		}
	}

	// attack error - opponent points
	// opponent attack - opponent points
	// serve error - opponent points
	// opponent serve - opponent points
	// opponent block - opponent points - attack neutrals
	// block error - opponent points - opponent attacks
	// error - opponent points
	if rally.Action == "attack_error" ||
		rally.Action == "opponent_attack" ||
		rally.Action == "serve_error" ||
		rally.Action == "opponent_serve" ||
		rally.Action == "opponent_block" ||
		rally.Action == "block_error" ||
		rally.Action == "error" {
		query := "UPDATE set SET opponent_points = opponent_points + 1 WHERE set_id = $1"
		result, err := tx.ExecContext(
			ctx,
			query,
			rally.SetId,
		)
		if err != nil {
			return fail(err)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil || int(rowsAffected) != 1 {
			return fail(err)
		}

		if rally.Action == "opponent_block" {
			query := "UPDATE set SET attack_neutrals = attack_neutrals + 1 WHERE set_id = $1"
			result, err := tx.ExecContext(
				ctx,
				query,
				rally.SetId,
			)
			if err != nil {
				return fail(err)
			}
			rowsAffected, err := result.RowsAffected()
			if err != nil || int(rowsAffected) != 1 {
				return fail(err)
			}
		}

		if rally.Action == "block_error" {
			query := "UPDATE set SET opponent_attacks = opponent_attacks + 1 WHERE set_id = $1"
			result, err := tx.ExecContext(
				ctx,
				query,
				rally.SetId,
			)
			if err != nil {
				return fail(err)
			}
			rowsAffected, err := result.RowsAffected()
			if err != nil || int(rowsAffected) != 1 {
				return fail(err)
			}
		}
	}
	// Commit transaction
	if err := tx.Commit(); err != nil {
		return fail(err)
	}
	return int(rowsAffected), nil
}
