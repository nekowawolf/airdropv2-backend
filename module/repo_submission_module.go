package module

import (
	"fmt"
	"time"

	"github.com/nekowawolf/airdropv2/config"
	"github.com/nekowawolf/airdropv2/models"
	"github.com/nekowawolf/airdropv2/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRepoSubmission(repoURL string, addedBy *models.AddedByInfo) interface{} {
	newSubmission := models.RepoSubmission{
		ID:        primitive.NewObjectID(),
		RepoURL:   repoURL,
		AddedBy:   addedBy,
		CreatedAt: time.Now(),
	}

	insertedID, err := InsertDocument("repoSubmissions", newSubmission)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return insertedID
}

func GetAllRepoSubmissions() ([]models.RepoSubmission, error) {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	collection := config.Database.Collection("repoSubmissions")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error retrieving data: %v", err)
	}
	defer cursor.Close(ctx)

	var submissions []models.RepoSubmission
	if err = cursor.All(ctx, &submissions); err != nil {
		return nil, fmt.Errorf("error decoding data: %v", err)
	}

	return submissions, nil
}

func DeleteRepoSubmissionByID(id primitive.ObjectID) error {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	collection := config.Database.Collection("repoSubmissions")
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("error deleting repo submission for ID %s: %s", id.Hex(), err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no repo submission found with ID %s", id.Hex())
	}

	return nil
}