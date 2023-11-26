package domain

import (
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type TeamMainInfo struct {
	TeamId         int       `json:"teamId"`
	UserId         int       `json:"userId"`
	Name           string    `json:"name" binding:"required,max=30"`
	Category       string    `json:"category" binding:"required,teamcategory"`
	Country        string    `json:"country" binding:"required,min=4"`
	Province       string    `json:"province" binding:"required"`
	City           string    `json:"city" binding:"required"`
	CreationDate   time.Time `json:"creationDateTime"`
	LastUpdateDate time.Time `json:"lastUpdateDateTime"`
}

type TeamGameData struct {
	WonGames            int     `json:"wonGames"`
	TotalGames          int     `json:"totalGames"`
	WonSets             int     `json:"teamSets"`
	TotalSets           int     `json:"totalSets"`
	AttackPoints        int     `json:"attackPoints"`
	AttackNeutrals      int     `json:"attackNeutrals"`
	AttackErrors        int     `json:"attackErrors"`
	TotalAttacks        int     `json:"totalAttacks"`
	AttackEffectiveness float64 `json:"attackEffectiveness"`
	BlockPoints         int     `json:"blockPoints"`
	BlockNeutrals       int     `json:"blockNeutrals"`
	BlockErrors         int     `json:"blockErrors"`
	TotalBlocks         int     `json:"totalBlocks"`
	BlockEffectiveness  float64 `json:"blockEffectiveness"`
	ServePoints         int     `json:"servePoints"`
	ServeNeutrals       int     `json:"serveNeutrals"`
	ServeErrors         int     `json:"serveErrors"`
	TotalServes         int     `json:"totalServes"`
	ServeEffectiveness  float64 `json:"serveEffectiveness"`
	OpponentErrors      int     `json:"opponentErrors"`
	TotalPoints         int     `json:"totalPoints"`
	TotalActions        int     `json:"totalActions"`
	TotalEffectiveness  float64 `json:"totalEffectiveness"`
	OpponentAttacks     int     `json:"opponentAttacks"`
	OpponentBlocks      int     `json:"opponentBlocks"`
	OpponentServes      int     `json:"opponentServes"`
	TotalErrors         int     `json:"totalErrors"`
}

type Team struct {
	TeamMainInfo TeamMainInfo `json:"teamMainInfo"`
	TeamGameData TeamGameData `json:"teamGameData"`
}

var ValidTeamCategory validator.Func = func(fl validator.FieldLevel) bool {
	category, ok := fl.Field().Interface().(string)
	if ok {
		if category == "Women" || category == "Men" {
			return true
		}
	}
	return false
}

// Registers custom validators in models for JSON binding
func RegisterTeamValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("teamcategory", ValidTeamCategory)
	}
}

// RESPONSE STRUCTS

type TeamSummary struct {
	TeamMainInfo        TeamMainInfo `json:"teamMainInfo"`
	WonGames            int          `json:"wonGames"`
	TotalGames          int          `json:"totalGames"`
	WonSets             int          `json:"teamSets"`
	TotalSets           int          `json:"totalSets"`
	AttackEffectiveness float64      `json:"attackEffectiveness"`
	BlockEffectiveness  float64      `json:"blockEffectiveness"`
	ServeEffectiveness  float64      `json:"serveEffectiveness"`
	TotalEffectiveness  float64      `json:"totalEffectiveness"`
}
