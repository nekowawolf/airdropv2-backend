package portfolio

import (
	"errors"
	"fmt"

	"github.com/nekowawolf/airdropv2/config"

	"github.com/nekowawolf/airdropv2/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const projectsCollection = "portfolio_projects"

func InsertProject(req Project) (interface{}, error) {
	req.ID = primitive.NewObjectID()
	return utils.InsertDocument(projectsCollection, req)
}

func GetAllProjects() ([]Project, error) {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	collection := config.Database.Collection(projectsCollection)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("GetAllProjects Find: %v", err)
	}

	var items []Project
	if err = cursor.All(ctx, &items); err != nil {
		return nil, fmt.Errorf("GetAllProjects All: %v", err)
	}

	return items, nil
}

func GetProjectByID(id primitive.ObjectID) (Project, error) {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	collection := config.Database.Collection(projectsCollection)
	var item Project
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&item)
	return item, err
}

func UpdateProjectByID(id primitive.ObjectID, updateData Project) error {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"title":       updateData.Title,
			"description": updateData.Description,
			"image_url":   updateData.ImageURL,
			"link":        updateData.Link,
			"github_url":  updateData.GitHubURL,
			"screenshots": updateData.Screenshots,
			"ss_desc":     updateData.SSDesc,
			"video_url":   updateData.VideoURL,
			"use_case":    updateData.UseCase,
			"activity":    updateData.Activity,
			"erd":         updateData.ERD,
			"flowchart":   updateData.Flowchart,
			"stack":       updateData.Stack,
		},
	}

	result, err := config.Database.Collection(projectsCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("UpdateProjectByID: %v", err)
	}

	if result.ModifiedCount == 0 {
		return errors.New("no data has been changed with the specified ID")
	}

	return nil
}

func DeleteProjectByID(id primitive.ObjectID) error {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	collection := config.Database.Collection(projectsCollection)
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("error deleting project ID %s: %s", id.Hex(), err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("project with ID %s not found", id.Hex())
	}

	return nil
}