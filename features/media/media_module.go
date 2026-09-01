package media

import (
	"fmt"
	"time"

	"github.com/nekowawolf/airdropv2/config"

	"github.com/nekowawolf/airdropv2/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertMedia(filename, url string, size int64, contentType, mediaType, r2Key string) interface{} {
	newMedia := Media{
		ID:          primitive.NewObjectID(),
		Filename:    filename,
		URL:         url,
		Size:        size,
		ContentType: contentType,
		MediaType:   mediaType,
		R2Key:       r2Key,
		CreatedAt:   time.Now(),
	}

	insertedID, err := utils.InsertDocument("media", newMedia)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Printf("Inserted new media with ID: %v\n", insertedID)
	return insertedID
}

func GetAllMedia() ([]Media, error) {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	collection := config.Database.Collection("media")

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, fmt.Errorf("error retrieving data: %v", err)
	}
	defer cursor.Close(ctx)

	var mediaList []Media
	if err = cursor.All(ctx, &mediaList); err != nil {
		return nil, fmt.Errorf("error decoding data: %v", err)
	}

	return mediaList, nil
}

func GetMediaByID(id primitive.ObjectID) (*Media, error) {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	collection := config.Database.Collection("media")
	filter := bson.M{"_id": id}

	var result Media
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func DeleteMediaByID(id primitive.ObjectID) error {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	collection := config.Database.Collection("media")
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("error deleting media for ID %s: %s", id.Hex(), err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no media found with ID %s", id.Hex())
	}

	return nil
}