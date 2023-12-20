package domain

import (
	"time"
)

type GameMainInfo struct {
	TeamId         int       `json:"teamId"`
	GameDate       time.Time `json:"gameDate"`
	IsActive       bool      `json:"isActive"`
	GameCountry    string    `json:"gameCountry"`
	GameProvince   string    `json:"gameProvince"`
	GameCity       string    `json:"gameCity"`
	Opponent       string    `json:"opponent"`
	TeamSets       int       `json:"teamSets"`
	OpponentSets   int       `json:"opponentSets"`
	LastUpdateDate time.Time `json:"lastUpdateDate"`
}

type GameTeamsNames struct {
	TeamName     string `json:"teamName"`
	OpponentName string `json:"opponentName"`
}

type GameSummary struct {
	GameId          int       `json:"gameId"`
	TeamId          int       `json:"teamId"`
	GameDate        time.Time `json:"gameDate"`
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
	GameWinner      string    `json:"gameWinner"`
	LastUpdateDate  time.Time `json:"lastUpdateDate"`
}

type Game struct {
	GameId              int       `json:"gameId"`
	TeamId              int       `json:"teamId"`
	GameDate            time.Time `json:"gameDate"`
	IsActive            bool      `json:"isActive"`
	GameCountry         string    `json:"gameCountry"`
	GameProvince        string    `json:"gameProvince"`
	GameCity            string    `json:"gameCity"`
	Opponent            string    `json:"gameOpponent"`
	TeamSets            int       `json:"teamSets"`
	OpponentSets        int       `json:"opponentSets"`
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
	TotalErrors         int       `json:"totalErrors"`
	OpponentPoints      int       `json:"opponentPoints"`
	GameWinner          string    `json:"gameWinner"`
	LastUpdateDate      time.Time `json:"lastUpdateDate"`
}
