package repositories

import (
	"database/sql"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
)

type TeamRepository struct {
	db *sql.DB
}

var _ ports.TeamRepository = (*TeamRepository)(nil)

func NewTeamRepository(db *sql.DB) *TeamRepository {
	return &TeamRepository{
		db: db,
	}
}

func (t *TeamRepository) CheckTeamExistence(email string) (bool, error) {
	// query := bson.D{bson.E{Key: "email", Value: email}}
	// teamsWithSameEmail, err := database.Collection.CountDocuments(t.context, query)
	// if err != nil {
	// 	config.Logger.Error("No connection with database. Method: TeamService/checkTeamExistence/CountDocuments")
	// 	return true, err
	// }
	// if teamsWithSameEmail > 0 {
	// 	return true, nil
	// }
	// return false, nil
	return true, nil
}

func (t *TeamRepository) CreateTeam(team domain.NewTeam) (bool, error) {
	// query := "INSERT into "
	// result, err := database.Collection.InsertOne(t.context, team)
	// if err != nil {
	// 	errorMsg := fmt.Errorf(
	// 		"DATABASE error CreateTeam/InsertOne: %v",
	// 		err,
	// 	)
	// 	return false, errorMsg
	// }
	// _, ok := result.InsertedID.(primitive.ObjectID)
	// if !ok {
	// 	errorMsg := fmt.Errorf(
	// 		"DATABASE error CreateTeam/InsertedID: %v",
	// 		err,
	// 	)
	// 	return false, errorMsg
	// }
	return true, nil
}

func (t *TeamRepository) GetTeam(teamId string) (domain.Team, error) {
	var team domain.Team
	// id, _ := primitive.ObjectIDFromHex(teamId)
	// query := bson.D{bson.E{Key: "_id", Value: id}}
	// err := database.Collection.FindOne(t.context, query).Decode(&team)
	// if err != nil {
	// 	errorMsg := fmt.Errorf(
	// 		"DATABASE error GetTeam/FindOne: %v",
	// 		err,
	// 	)
	// 	return team, errorMsg
	// }
	return team, nil
}

func (t *TeamRepository) UpdateTeamInfo(team domain.BaseTeam) (bool, error) {
	return true, nil
}
