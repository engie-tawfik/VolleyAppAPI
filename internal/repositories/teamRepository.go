package repositories

import (
	"context"
	"volleyapp/config"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
	"volleyapp/internal/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TeamRepository struct {
	collection *mongo.Collection
	context    context.Context
}

var _ ports.TeamRepository = (*TeamRepository)(nil)

func NewTeamRepository(collection *mongo.Collection, context context.Context) *TeamRepository {
	return &TeamRepository{
		collection: collection,
		context:    context,
	}
}

func (t *TeamRepository) CheckTeamExistence(email string) (bool, error) {
	query := bson.D{bson.E{Key: "email", Value: email}}
	teamsWithSameEmail, err := database.Collection.CountDocuments(t.context, query)
	if err != nil {
		config.Logger.Error("No connection with database. Method: TeamService/checkTeamExistence/CountDocuments")
		return true, err
	}
	if teamsWithSameEmail > 0 {
		return true, nil
	}
	return false, nil
}

func (t *TeamRepository) CreateTeam(team domain.NewTeam) (string, bool) {
	result, err := database.Collection.InsertOne(t.context, team)
	if err != nil {
		config.Logger.Error("No connection with database. Method: TeamService/CreateTeam/InsertOne")
		return "No connection with database.", false
	}
	teamId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		config.Logger.Error("Unable to create team")
		return "Unable to create team.", false
	}
	return teamId.Hex(), true
}

func (t *TeamRepository) GetTeam(teamId string) domain.Team {
	var team domain.Team
	id, _ := primitive.ObjectIDFromHex(teamId)
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := database.Collection.FindOne(t.context, query).Decode(&team)
	if err != nil {
		config.Logger.Error("No connection with database. Method: GetTeam/FindOne")
		return team
	}
	return team
}
