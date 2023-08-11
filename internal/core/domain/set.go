package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Set struct and methods
type Set struct {
	SetId               primitive.ObjectID `json:"setId" bson:"setId"`
	ActiveStatus        bool               `json:"activeStatus" bson:"activeStatus"`
	Opponent            string             `json:"gameOpponent" bson:"gameOpponent"`
	TeamSets            int                `json:"teamSets" bson:"teamSets"`
	OpponentSets        int                `json:"opponentSets" bson:"opponentSets"`
	Sets                []Set              `json:"gameSets" bson:"gameSets"`
	AttackPoints        int                `json:"attackPoints" bson:"attackPoints"`
	AttackNeutrals      int                `json:"attackNeutrals" bson:"attackNeutrals"`
	AttackErrors        int                `json:"attackErrors" bson:"attackErrors"`
	TotalAttacks        int                `json:"totalAttacks" bson:"totalAttacks"`
	AttackEffectiveness float64            `json:"attackEffectiveness" bson:"attackEffectiveness"`
	BlockPoints         int                `json:"blockPoints" bson:"blockPoints"`
	BlockNeutrals       int                `json:"blockNeutrals" bson:"blockNeutrals"`
	BlockErrors         int                `json:"blockErrors" bson:"blockErrors"`
	TotalBlocks         int                `json:"totalBlocks" bson:"totalBlocks"`
	BlockEffectiveness  float64            `json:"blockEffectiveness" bson:"blockEffectiveness"`
	ServePoints         int                `json:"servePoints" bson:"servePoints"`
	ServeNeutrals       int                `json:"serveNeutrals" bson:"serveNeutrals"`
	ServeErrors         int                `json:"serveErrors" bson:"serveErrors"`
	TotalServes         int                `json:"totalServes" bson:"totalServes"`
	ServeEffectiveness  float64            `json:"serveEffectiveness" bson:"serveEffectiveness"`
	OpponentErrors      int                `json:"opponentErrors" bson:"opponentErrors"`
	TotalPoints         int                `json:"totalPoints" bson:"totalPoints"`
	TotalActions        int                `json:"totalActions" bson:"totalActions"`
	TotalEffectiveness  float64            `json:"totalEffectiveness" bson:"totalEffectiveness"`
	OpponentAttacks     int                `json:"opponentAttacks" bson:"opponentAttacks"`
	OpponentBlocks      int                `json:"opponentBlocks" bson:"opponentBlocks"`
	OpponentServes      int                `json:"opponentServes" bson:"opponentServes"`
	TotalErrors         int                `json:"totalErrors" bson:"totalErrors"`
	OpponentPoints      int                `json:"opponentPoints" bson:"opponentPoints"`
	Winner              string             `json:"gameWinner" bson:"gameWinner"`
	GameActions         []string           `json:"gameActions" bson:"gameActions"`
	Forward             bool               `json:"forward" bson:"forward"`
	SetCount            int                `json:"setCount" bson:"setCount"`
}
