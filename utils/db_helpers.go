package utils

import (
	"fmt"
	"github.com/nekowawolf/airdropv2/config"
)

func InsertDocument(collection string, doc interface{}) (interface{}, error) {
	ctx, cancel := GetDBContext()
	defer cancel()

	insertResult, err := config.Database.Collection(collection).InsertOne(ctx, doc)
	if err != nil {
		return nil, fmt.Errorf("InsertDocument error in collection %s: %v", collection, err)
	}
	return insertResult.InsertedID, nil
}