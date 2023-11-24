package domain

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginTeam struct {
	TeamId   primitive.ObjectID `json:"teamId" bson:"_id"`
	Password string             `json:"password" bson:"password"`
	Email    string             `json:"email" bson:"email"`
}

type BaseTeam struct {
	Name               string    `json:"name" bson:"name" binding:"required,max=30"`
	Category           string    `json:"category" bson:"category" binding:"required,teamcategory"`
	Country            string    `json:"country" bson:"country" binding:"required,min=4"`
	Province           string    `json:"province" bson:"province" binding:"required"`
	City               string    `json:"city" bson:"city" binding:"required"`
	Email              string    `json:"email" bson:"email" binding:"required,email"`
	CreationDateTime   time.Time `json:"creationDateTime" bson:"creationDateTime"`
	LastUpdateDateTime time.Time `json:"lastUpdateDateTime" bson:"lastUpdateDateTime"`
}

type NewTeam struct {
	BaseTeam `bson:",inline"`
	Password string `json:"password" bson:"password" binding:"required,passwordcheck"`
	TeamData `bson:",inline"`
}

type TeamData struct {
	Games               []Game  `json:"teamGames" bson:"teamGames"`
	WonGames            int     `json:"wonGames" bson:"wonGames"`
	TotalGames          int     `json:"totalGames" bson:"totalGames"`
	WonSets             int     `json:"teamSets" bson:"teamSets"`
	TotalSets           int     `json:"totalSets" bson:"totalSets"`
	AttackPoints        int     `json:"attackPoints" bson:"attackPoints"`
	AttackNeutrals      int     `json:"attackNeutrals" bson:"attackNeutrals"`
	AttackErrors        int     `json:"attackErrors" bson:"attackErrors"`
	TotalAttacks        int     `json:"totalAttacks" bson:"totalAttacks"`
	AttackEffectiveness float64 `json:"attackEffectiveness" bson:"attackEffectiveness"`
	BlockPoints         int     `json:"blockPoints" bson:"blockPoints"`
	BlockNeutrals       int     `json:"blockNeutrals" bson:"blockNeutrals"`
	BlockErrors         int     `json:"blockErrors" bson:"blockErrors"`
	TotalBlocks         int     `json:"totalBlocks" bson:"totalBlocks"`
	BlockEffectiveness  float64 `json:"blockEffectiveness" bson:"blockEffectiveness"`
	ServePoints         int     `json:"servePoints" bson:"servePoints"`
	ServeNeutrals       int     `json:"serveNeutrals" bson:"serveNeutrals"`
	ServeErrors         int     `json:"serveErrors" bson:"serveErrors"`
	TotalServes         int     `json:"totalServes" bson:"totalServes"`
	ServeEffectiveness  float64 `json:"serveEffectiveness" bson:"serveEffectiveness"`
	OpponentErrors      int     `json:"opponentErrors" bson:"opponentErrors"`
	TotalPoints         int     `json:"totalPoints" bson:"totalPoints"`
	TotalActions        int     `json:"totalActions" bson:"totalActions"`
	TotalEffectiveness  float64 `json:"totalEffectiveness" bson:"totalEffectiveness"`
	OpponentAttacks     int     `json:"opponentAttacks" bson:"opponentAttacks"`
	OpponentBlocks      int     `json:"opponentBlocks" bson:"opponentBlocks"`
	OpponentServes      int     `json:"opponentServes" bson:"opponentServes"`
	TotalErrors         int     `json:"totalErrors" bson:"totalErrors"`
}

type Team struct {
	BaseTeam `bson:",inline"`
	TeamData `bson:",inline"`
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

func GetTeamErrorMsg(fe validator.FieldError) string {
	var msg string
	if fe.Tag() == "required" {
		msg = fmt.Sprintf("%s is missing in team's data.", fe.Field())
	} else if fe.Tag() == "min" || fe.Tag() == "max" || fe.Tag() == "email" || fe.Tag() == "teamcategory" || fe.Tag() == "passwordcheck" {
		msg = fmt.Sprintf("Invalid value for %s", fe.Field())
	}
	return msg
}

// Registers custom validators in models for JSON binding
func RegisterTeamValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("teamcategory", ValidTeamCategory)
		v.RegisterValidation("passwordcheck", PasswordCheck)
	}
}
