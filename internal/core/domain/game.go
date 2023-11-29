package domain

import (
	"time"
)

type Game struct {
	GameId              int       `json:"gameId"`
	TeamId              int       `json:"teamId"`
	GameDate            time.Time `json:"gameDateTime"`
	IsActive            bool      `json:"isActive"`
	GameCountry         string    `json:"gameCountry"`
	GameProvince        string    `json:"gameProvince"`
	GameCity            string    `json:"gameCity"`
	Opponent            string    `json:"gameOpponent"`
	TeamSets            int       `json:"teamSets"`
	OpponentSets        int       `json:"opponentSets"`
	AttackPoints        int       `json:"attackPoints"`
	AttackNeutrals      int       `json:"attackNeutrals"`
	AttackErrors        int       `json:"attackErrors"`
	TotalAttacks        int       `json:"totalAttacks"`
	AttackEffectiveness float64   `json:"attackEffectiveness"`
	BlockPoints         int       `json:"blockPoints"`
	BlockNeutrals       int       `json:"blockNeutrals"`
	BlockErrors         int       `json:"blockErrors"`
	TotalBlocks         int       `json:"totalBlocks"`
	BlockEffectiveness  float64   `json:"blockEffectiveness"`
	ServePoints         int       `json:"servePoints"`
	ServeNeutrals       int       `json:"serveNeutrals"`
	ServeErrors         int       `json:"serveErrors"`
	TotalServes         int       `json:"totalServes"`
	ServeEffectiveness  float64   `json:"serveEffectiveness"`
	OpponentErrors      int       `json:"opponentErrors"`
	TotalPoints         int       `json:"totalPoints"`
	TotalActions        int       `json:"totalActions"`
	TotalEffectiveness  float64   `json:"totalEffectiveness"`
	OpponentAttacks     int       `json:"opponentAttacks"`
	OpponentBlocks      int       `json:"opponentBlocks"`
	OpponentServes      int       `json:"opponentServes"`
	TotalErrors         int       `json:"totalErrors"`
	OpponentPoints      int       `json:"opponentPoints"`
	Winner              string    `json:"gameWinner"`
	LastUpdateDate      time.Time `json:"lastUpdateDate"`
}

type GameMainInfo struct {
	TeamId         int       `json:"teamId"`
	GameDate       time.Time `json:"gameDateTime"`
	IsActive       bool      `json:"isActive"`
	GameCountry    string    `json:"gameCountry"`
	GameProvince   string    `json:"gameProvince"`
	GameCity       string    `json:"gameCity"`
	Opponent       string    `json:"gameOpponent"`
	LastUpdateDate time.Time `json:"lastUpdateDate"`
}

type GameSummary struct {
	GameId          int       `json:"gameId"`
	TeamId          int       `json:"teamId"`
	GameDateTime    time.Time `json:"gameDateTime"`
	IsActive        bool      `json:"isActive"`
	GameCountry     string    `json:"gameCountry"`
	GameProvince    string    `json:"gameProvince"`
	GameCity        string    `json:"gameCity"`
	Opponent        string    `json:"gameOpponent"`
	TeamSets        int       `json:"teamSets"`
	OpponentSets    int       `json:"opponentSets"`
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
	Winner          string    `json:"gameWinner"`
	LastUpdateDate  time.Time `json:"lastUpdateDate"`
}
