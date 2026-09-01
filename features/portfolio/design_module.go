package portfolio

import (
	"errors"
	"fmt"

	"github.com/nekowawolf/airdropv2/config"

	"github.com/nekowawolf/airdropv2/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const designsCollection = "portfolio_designs"

func InsertDesign(req Design) (interface{}, error) {
	req.ID = primitive.NewObjectID()
	return utils.InsertDocument(designsCollection, req)
}

func GetAllDesigns() ([]Design, error) {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	collection := config.Database.Collection(designsCollection)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("GetAllDesigns Find: %v", err)
	}

	var items []Design
	if err = cursor.All(ctx, &items); err != nil {
		return nil, fmt.Errorf("GetAllDesigns All: %v", err)
	}

	return items, nil
}

func GetDesignByID(id primitive.ObjectID) (Design, error) {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	collection := config.Database.Collection(designsCollection)
	var item Design
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&item)
	return item, err
}

func UpdateDesignByID(id primitive.ObjectID, updateData Design) error {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"title":         updateData.Title,
			"description":   updateData.Description,
			"image_url":     updateData.ImageURL,
			"link":          updateData.Link,
			"video_url":     updateData.VideoURL,
			"category":      updateData.Category,
			"tools":         updateData.Tools,
			"screenshots":   updateData.Screenshots,
			"ss_desc":       updateData.SSDesc,
			"color_palette": updateData.ColorPalette,
			"typography":    updateData.Typography,
		},
	}

	result, err := config.Database.Collection(designsCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("UpdateDesignByID: %v", err)
	}

	if result.ModifiedCount == 0 {
		return errors.New("no data has been changed with the specified ID")
	}

	return nil
}

func DeleteDesignByID(id primitive.ObjectID) error {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	collection := config.Database.Collection(designsCollection)
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("error deleting design ID %s: %s", id.Hex(), err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("design with ID %s not found", id.Hex())
	}

	return nil
}