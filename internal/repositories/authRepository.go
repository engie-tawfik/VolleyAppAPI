package repositories

import (
	"context"
	"log"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
	"volleyapp/internal/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	collection *mongo.Collection
	context    context.Context
}

var _ ports.AuthRepository = (*AuthRepository)(nil)

func NewAuthRepository(collection *mongo.Collection, context context.Context) *AuthRepository {
	return &AuthRepository{
		collection: collection,
		context:    context,
	}
}

func (a *AuthRepository) Login(email string) domain.LoginTeam {
	// Get team based on email
	var team domain.LoginTeam
	query := bson.D{bson.E{Key: "email", Value: email}}
	err := database.Collection.FindOne(a.context, query).Decode(&team)
	if err != nil {
		log.Println("Email not found. Error:", err)
		return team
	}
	return team
}
