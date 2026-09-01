package portfolio

import (
	"errors"
	"fmt"

	"github.com/nekowawolf/airdropv2/config"

	"github.com/nekowawolf/airdropv2/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const certificatesCollection = "portfolio_certificates"

func InsertCertificate(req Certificate) (interface{}, error) {
	req.ID = primitive.NewObjectID()
	return utils.InsertDocument(certificatesCollection, req)
}

func GetAllCertificates() ([]Certificate, error) {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	collection := config.Database.Collection(certificatesCollection)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("GetAllCertificates Find: %v", err)
	}

	var items []Certificate
	if err = cursor.All(ctx, &items); err != nil {
		return nil, fmt.Errorf("GetAllCertificates All: %v", err)
	}

	return items, nil
}

func GetCertificateByID(id primitive.ObjectID) (Certificate, error) {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	collection := config.Database.Collection(certificatesCollection)
	var item Certificate
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&item)
	return item, err
}

func UpdateCertificateByID(id primitive.ObjectID, updateData Certificate) error {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"title":     updateData.Title,
			"image_url": updateData.ImageURL,
		},
	}

	result, err := config.Database.Collection(certificatesCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("UpdateCertificateByID: %v", err)
	}

	if result.ModifiedCount == 0 {
		return errors.New("no data has been changed with the specified ID")
	}

	return nil
}

func DeleteCertificateByID(id primitive.ObjectID) error {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	collection := config.Database.Collection(certificatesCollection)
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("error deleting certificate ID %s: %s", id.Hex(), err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("certificate with ID %s not found", id.Hex())
	}

	return nil
}