package domain

import "time"

type Set struct {
	SetId               int       `json:"setId"`
	GameId              int       `json:"gameId"`
	StartedAt           time.Time `json:"startedAt"`
	IsActive            bool      `json:"isActive"`
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
	SetWinner           string    `json:"setWinner"`
	GameActions         []string  `json:"gameActions"`
	Forward             bool      `json:"forward"`
	SetCount            int       `json:"setCount"`
	LastUpdate          time.Time `json:"lastUpdate"`
}

type SetMainInfo struct {
	GameId     int       `json:"gameId"`
	StartedAt  time.Time `json:"startedAt"`
	IsActive   bool      `json:"isActive"`
	LastUpdate time.Time `json:"lastUpdate"`
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
