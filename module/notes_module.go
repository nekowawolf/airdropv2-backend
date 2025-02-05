package module

import (
	"context"
	"time"
	"fmt"
	"github.com/nekowawolf/airdropv2/config"
	"github.com/nekowawolf/airdropv2/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertOneDocNotes(collection string, doc interface{}) (interface{}, error) {
	insertResult, err := config.Database.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %v", err)
	}
	return insertResult.InsertedID, nil
}

func InsertNotes(title, content string) (interface{}, error) {
	notes := models.Notes{
		ID:      primitive.NewObjectID(),
		Title:   title,
		Content: content,
		Date:    time.Now(),
	}

	insertedID, err := InsertOneDocNotes("notes", notes)
	if err != nil {
		return nil, err 
	}

	fmt.Printf("Inserted new note with ID: %v\n", insertedID)
	return insertedID, nil
}

func GetAllNotes() ([]models.Notes, error) {
	collection := config.Database.Collection("notes")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("GetAllNotes Find: %v", err)
	}

	var notes []models.Notes
	if err = cursor.All(context.TODO(), &notes); err != nil {
		return nil, fmt.Errorf("GetAllNotes All: %v", err)
	}

	return notes, nil
}

func GetNotesByID(id primitive.ObjectID) (models.Notes, error) {
	collection := config.Database.Collection("notes")
	var notes models.Notes
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&notes)
	if err != nil {
		return models.Notes{}, err
	}
	return notes, nil
}

func UpdateNotesByID(id primitive.ObjectID, title string, content string) (models.Notes, error) {
	collection := config.Database.Collection("notes")

	update := bson.M{
		"$set": bson.M{
			"title":   title,
			"content": content,
		},
	}

	filter := bson.M{"_id": id}

	var updatedNote models.Notes
	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&updatedNote)
	if err != nil {
		return models.Notes{}, err
	}

	return updatedNote, nil
}

func DeleteNotesByID(id primitive.ObjectID) error {
    collection := config.Database.Collection("notes")
    filter := bson.M{"_id": id}

    result, err := collection.DeleteOne(context.TODO(), filter)
    if err != nil {
        return fmt.Errorf("error deleting note for ID %s: %s", id.Hex(), err.Error())
    }

    if result.DeletedCount == 0 {
        return fmt.Errorf("no note found with ID %s", id.Hex())
    }

    return nil
}