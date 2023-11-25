package domain

import (
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginTeam struct {
	TeamId   primitive.ObjectID `json:"teamId"`
	Password string             `json:"password"`
	Email    string             `json:"email"`
}

type BaseTeam struct {
	Name               string    `json:"name" binding:"required,max=30"`
	Category           string    `json:"category" binding:"required,teamcategory"`
	Country            string    `json:"country" binding:"required,min=4"`
	Province           string    `json:"province" binding:"required"`
	City               string    `json:"city" binding:"required"`
	Email              string    `json:"email" binding:"required,email"`
	CreationDateTime   time.Time `json:"creationDateTime"`
	LastUpdateDateTime time.Time `json:"lastUpdateDateTime"`
}

type NewTeam struct {
	BaseTeam
	Password string `json:"password" binding:"required,passwordcheck"`
	TeamData
}

type TeamData struct {
	Games               []Game  `json:"teamGames"`
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
	BaseTeam
	TeamData
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
		v.RegisterValidation("passwordcheck", PasswordCheck)
	}
}
