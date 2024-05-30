package models

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	Refreshtoken string `json:"refreshToken"`
}

type Auth struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
