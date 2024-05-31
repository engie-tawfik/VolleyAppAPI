package models

import (
	"math"
	"time"
	"volleyapp/domain/constants"
	"volleyapp/utils"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type SetMainInfo struct {
	GameId         int       `json:"gameId"`
	StartedAt      time.Time `json:"startedAt"`
	IsActive       bool      `json:"isActive"`
	TotalPoints    int       `json:"totalPoints"`
	OpponentPoints int       `json:"opponentPoints"`
	SetWinner      string    `json:"setWinner"`
	LastUpdate     time.Time `json:"lastUpdate"`
	GameOpponent   string    `json:"gameOpponent"`
	TeamId         int       `json:"teamId"`
	TeamName       string    `json:"teamName"`
}

type SetSummary struct {
	SetId           int       `json:"setId"`
	GameId          int       `json:"gameId"`
	StartedAt       time.Time `json:"startedAt"`
	IsActive        bool      `json:"isActive"`
	AttackPoints    int       `json:"attackPoints"`
	OpponentAttacks int       `json:"opponentAttacks"`
	BlockPoints     int       `json:"blockPoints"`
	OpponentBlocks  int       `json:"opponentBlocks"`
	ServePoints     int       `json:"servePoints"`
	OpponentServes  int       `json:"opponentServes"`
	TotalErrors     int       `json:"totalErrors"`
	OpponentErrors  int       `json:"opponentErrors"`
	AttackErrors    int       `json:"attackErrors"`
	ServeErrors     int       `json:"serveErrors"`
	TotalPoints     int       `json:"totalPoints"`
	OpponentPoints  int       `json:"opponentPoints"`
	SetWinner       string    `json:"setWinner"`
	LastUpdate      time.Time `json:"lastUpdate"`
}

type Rally struct {
	SetId  int    `json:"setId" binding:"required"`
	Action string `json:"action" binding:"required,gameactions"`
}

type Set struct {
	SetId               int       `json:"setId"`
	GameId              int       `json:"gameId"`
	StartedAt           time.Time `json:"startedAt"`
	IsActive            bool      `json:"isActive"`
	TotalAttacks        int       `json:"totalAttacks"`
	AttackPoints        int       `json:"attackPoints"`
	AttackNeutrals      int       `json:"attackNeutrals"`
	AttackErrors        int       `json:"attackErrors"`
	AttackEffectiveness float64   `json:"attackEffectiveness"`
	TotalBlocks         int       `json:"totalBlocks"`
	BlockPoints         int       `json:"blockPoints"`
	BlockNeutrals       int       `json:"blockNeutrals"`
	BlockErrors         int       `json:"blockErrors"`
	BlockEffectiveness  float64   `json:"blockEffectiveness"`
	TotalServes         int       `json:"totalServes"`
	ServePoints         int       `json:"servePoints"`
	ServeNeutrals       int       `json:"serveNeutrals"`
	ServeErrors         int       `json:"serveErrors"`
	ServeEffectiveness  float64   `json:"serveEffectiveness"`
	OpponentErrors      int       `json:"opponentErrors"`
	TotalPoints         int       `json:"totalPoints"`
	TotalActions        int       `json:"totalActions"`
	TotalEffectiveness  float64   `json:"totalEffectiveness"`
	OpponentAttacks     int       `json:"opponentAttacks"`
	OpponentBlocks      int       `json:"opponentBlocks"`
	OpponentServes      int       `json:"opponentServes"`
	TotalErrors         int       `json:"errors"`
	OpponentPoints      int       `json:"opponentPoints"`
	SetWinner           string    `json:"setWinner"`
	GameActions         []string  `json:"gameActions"`
	Forward             bool      `json:"forward"`
	SetCount            int       `json:"setCount"`
	LastUpdate          time.Time `json:"lastUpdate"`
}

var ValidGameAction validator.Func = func(fl validator.FieldLevel) bool {
	action, ok := fl.Field().Interface().(string)
	if ok {
		if utils.CheckStringInArray(action, constants.SetActions) {
			return true
		}
	}
	return false
}

// Registers custom validators in models for JSON binding
func RegisterSetValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("gameactions", ValidGameAction)
	}
}

func (s *Set) AttackPoint(fwd bool) {
	if fwd {
		s.AttackPoints += 1
		s.TotalPoints += 1
	} else if !fwd {
		s.AttackPoints -= 1
		s.TotalPoints -= 1
	}
}

func (s *Set) AttackNeutral(fwd bool) {
	if fwd {
		s.AttackNeutrals += 1
	} else if !fwd {
		s.AttackNeutrals -= 1
	}
}

func (s *Set) AttackError(fwd bool) {
	if fwd {
		s.AttackErrors += 1
		s.OpponentPoints += 1
	} else if !fwd {
		s.AttackErrors -= 1
		s.OpponentPoints -= 1
	}
}

func (s *Set) OpponentAttack(fwd bool) {
	if fwd {
		s.OpponentAttacks += 1
		s.OpponentPoints += 1
	} else {
		s.OpponentAttacks -= 1
		s.OpponentPoints -= 1
	}
}

func (s *Set) BlockPoint(fwd bool) {
	if fwd {
		s.BlockPoints += 1
		s.TotalPoints += 1
	} else {
		s.BlockPoints -= 1
		s.TotalPoints -= 1
	}
}

func (s *Set) BlockNeutral(fwd bool) {
	if fwd {
		s.BlockNeutrals += 1
		s.TotalBlocks += 1
	} else {
		s.BlockNeutrals -= 1
		s.TotalBlocks -= 1
	}
}

func (s *Set) BlockError(fwd bool) {
	if fwd {
		s.BlockErrors += 1
		s.OpponentAttacks += 1
		s.OpponentPoints += 1
	} else {
		s.BlockErrors -= 1
		s.OpponentAttacks -= 1
		s.OpponentPoints -= 1
	}
}

func (s *Set) OpponentBlock(fwd bool) {
	if fwd {
		s.OpponentBlocks += 1
		s.OpponentPoints += 1
		s.AttackNeutrals += 1
	} else {
		s.OpponentBlocks -= 1
		s.OpponentPoints -= 1
		s.AttackNeutrals -= 1
	}
}

func (s *Set) ServePoint(fwd bool) {
	if fwd {
		s.ServePoints += 1
		s.TotalPoints += 1
	} else {
		s.ServePoints -= 1
		s.TotalPoints -= 1
	}
}

func (s *Set) ServeNeutral(fwd bool) {
	if fwd {
		s.ServeNeutrals += 1
	} else {
		s.ServeNeutrals -= 1
	}
}

func (s *Set) ServeError(fwd bool) {
	if fwd {
		s.ServeErrors += 1
		s.OpponentPoints += 1
	} else {
		s.ServeErrors -= 1
		s.OpponentPoints -= 1
	}
}

func (s *Set) OpponentServe(fwd bool) {
	if fwd {
		s.OpponentServes += 1
		s.OpponentPoints += 1
	} else {
		s.OpponentServes -= 1
		s.OpponentPoints -= 1
	}
}

func (s *Set) OpponentError(fwd bool) {
	if fwd {
		s.OpponentErrors += 1
		s.TotalPoints += 1
	} else {
		s.OpponentErrors -= 1
		s.TotalPoints -= 1
	}
}

func (s *Set) Error(fwd bool) {
	if fwd {
		s.TotalErrors += 1
		s.OpponentPoints += 1
	} else {
		s.TotalErrors -= 1
		s.OpponentPoints -= 1
	}
}

func (s *Set) UpdateStats() {
	// When doing rollback until set starting point (all stats in 0) effectiveness values are NaN
	// Ifs in this method fix it
	s.TotalAttacks = s.AttackPoints + s.AttackNeutrals + s.AttackErrors
	s.AttackEffectiveness =
		(float64(s.AttackPoints) / float64(s.TotalAttacks)) * 100
	if math.IsNaN(s.AttackEffectiveness) {
		s.AttackEffectiveness = 0.00
	}
	s.TotalBlocks = s.BlockPoints + s.BlockNeutrals + s.BlockErrors
	s.BlockEffectiveness =
		(float64(s.BlockPoints) / float64(s.TotalBlocks)) * 100
	if math.IsNaN(s.BlockEffectiveness) {
		s.BlockEffectiveness = 0.00
	}
	s.TotalServes = s.ServePoints + s.ServeNeutrals + s.ServeErrors
	s.ServeEffectiveness =
		(float64(s.ServePoints) / float64(s.TotalServes)) * 100
	if math.IsNaN(s.ServeEffectiveness) {
		s.ServeEffectiveness = 0.00
	}
	s.TotalActions =
		s.TotalAttacks + s.TotalBlocks + s.TotalServes + s.TotalErrors
	s.TotalEffectiveness =
		(float64(s.TotalPoints-s.OpponentErrors) / float64(s.TotalActions)) * 100
	if math.IsNaN(s.TotalEffectiveness) {
		s.TotalEffectiveness = 0.00
	}
}
