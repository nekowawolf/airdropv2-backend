package module

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/nekowawolf/airdropv2/config"
	"github.com/nekowawolf/airdropv2/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertOneDocAirdrop(collection string, doc interface{}) (interface{}, error) {
	collectionRef := config.Database.Collection(collection)
	insertResult, err := collectionRef.InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, fmt.Errorf("InsertOneDocAirdrop: %v", err)
	}
	return insertResult.InsertedID, nil
}

func InsertAirdropFree(name, task, link, level, status, backed, funds, supply, fdv, marketCap, vesting, linkClaim, linkDiscord, linkTwitter, imageURL, description, linkGuide string, price float64, usdIncome int) (interface{}, error) {
	var endedAt *time.Time
	if status == "ended" {
		now := time.Now()
		endedAt = &now
	}

	freeAirdrop := models.AirdropFree{
		ID:          primitive.NewObjectID(),
		Name:        name,
		Task:        task,
		Link:        link,
		Level:       level,
		Status:      status,
		Backed:      backed,
		Funds:       funds,
		Supply:      supply,
		Fdv:         fdv,
		MarketCap:   marketCap,
		Vesting:     vesting,
		LinkClaim:   linkClaim,
		LinkDiscord: linkDiscord,
		LinkTwitter: linkTwitter,
		ImageURL:    imageURL,
		Description: description,
		LinkGuide:   linkGuide,
		Price:       price,
		USDIncome:   usdIncome,
		CreatedAt:   time.Now(),
		EndedAt:     endedAt,
	}
	return InsertOneDocAirdrop("airdrop_free", freeAirdrop)
}

func InsertAirdropPaid(name, task, link, level, status, backed, funds, supply, fdv, marketCap, vesting, linkClaim, linkDiscord, linkTwitter, imageURL, description, linkGuide string, price float64, usdIncome int) (interface{}, error) {
	var endedAt *time.Time
	if status == "ended" {
		now := time.Now()
		endedAt = &now
	}

	paidAirdrop := models.AirdropPaid{
		ID:          primitive.NewObjectID(),
		Name:        name,
		Task:        task,
		Link:        link,
		Level:       level,
		Status:      status,
		Backed:      backed,
		Funds:       funds,
		Supply:      supply,
		Fdv:         fdv,
		MarketCap:   marketCap,
		Vesting:     vesting,
		LinkClaim:   linkClaim,
		LinkDiscord: linkDiscord,
		LinkTwitter: linkTwitter,
		ImageURL:    imageURL,
		Description: description,
		LinkGuide:   linkGuide,
		Price:       price,
		USDIncome:   usdIncome,
		CreatedAt:   time.Now(),
		EndedAt:     endedAt,
	}
	return InsertOneDocAirdrop("airdrop_paid", paidAirdrop)
}

func GetAllAirdrop() ([]interface{}, error) {
	var allAirdrops []interface{}

	freeAirdrops, err := GetAllAirdropFree()
	if err != nil {
		return nil, err
	}
	for _, free := range freeAirdrops {
		allAirdrops = append(allAirdrops, free)
	}

	paidAirdrops, err := GetAllAirdropPaid()
	if err != nil {
		return nil, err
	}
	for _, paid := range paidAirdrops {
		allAirdrops = append(allAirdrops, paid)
	}

	return allAirdrops, nil
}

func GetAllAirdropFree() ([]models.AirdropFree, error) {
	collection := config.Database.Collection("airdrop_free")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("GetAllAirdropFree Find: %v", err)
	}
	var airdrops []models.AirdropFree
	if err = cursor.All(context.TODO(), &airdrops); err != nil {
		return nil, fmt.Errorf("GetAllAirdropFree All: %v", err)
	}
	return airdrops, nil
}

func GetAllAirdropPaid() ([]models.AirdropPaid, error) {
	collection := config.Database.Collection("airdrop_paid")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("GetAllAirdropPaid Find: %v", err)
	}
	var airdrops []models.AirdropPaid
	if err = cursor.All(context.TODO(), &airdrops); err != nil {
		return nil, fmt.Errorf("GetAllAirdropPaid All: %v", err)
	}
	return airdrops, nil
}

func GetAllAirdropByID(id primitive.ObjectID) (interface{}, error) {
	freeAirdrop, err := GetAirdropFreeByID(id)
	if err == nil {
		return freeAirdrop, nil
	}

	paidAirdrop, err := GetAirdropPaidByID(id)
	if err == nil {
		return paidAirdrop, nil
	}

	return nil, fmt.Errorf("GetAllAirdropByID: airdrop not found in both collections")
}

func GetAirdropFreeByID(id primitive.ObjectID) (models.AirdropFree, error) {
	collection := config.Database.Collection("airdrop_free")
	var airdrop models.AirdropFree
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&airdrop)
	if err != nil {
		return models.AirdropFree{}, err
	}
	return airdrop, nil
}

func GetAirdropPaidByID(id primitive.ObjectID) (models.AirdropPaid, error) {
	collection := config.Database.Collection("airdrop_paid")
	var airdrop models.AirdropPaid
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&airdrop)
	if err != nil {
		return models.AirdropPaid{}, err
	}
	return airdrop, nil
}

func GetAirdropFreeByName(name string) ([]models.AirdropFree, error) {
	collection := config.Database.Collection("airdrop_free")
	filter := bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: name, Options: "i"}}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("GetAirdropFreeByName Find: %v", err)
	}
	var airdrops []models.AirdropFree
	if err = cursor.All(context.TODO(), &airdrops); err != nil {
		return nil, fmt.Errorf("GetAirdropFreeByName All: %v", err)
	}
	return airdrops, nil
}

func GetAirdropPaidByName(name string) ([]models.AirdropPaid, error) {
	collection := config.Database.Collection("airdrop_paid")
	filter := bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: name, Options: "i"}}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("GetAirdropPaidByName Find: %v", err)
	}
	var airdrops []models.AirdropPaid
	if err = cursor.All(context.TODO(), &airdrops); err != nil {
		return nil, fmt.Errorf("GetAirdropPaidByName All: %v", err)
	}
	return airdrops, nil
}

func GetAllAirdropByName(name string) ([]interface{}, error) {
	var allAirdrops []interface{}

	freeAirdrops, err := GetAirdropFreeByName(name)
	if err != nil {
		return nil, err
	}
	for _, free := range freeAirdrops {
		allAirdrops = append(allAirdrops, free)
	}

	paidAirdrops, err := GetAirdropPaidByName(name)
	if err != nil {
		return nil, err
	}
	for _, paid := range paidAirdrops {
		allAirdrops = append(allAirdrops, paid)
	}

	return allAirdrops, nil
}

func UpdateAllAirdropByID(id primitive.ObjectID, name, task, link, level, status, backed, funds, supply, fdv, marketCap, vesting, linkClaim, linkDiscord, linkTwitter, imageURL, description, linkGuide string, price float64, usdIncome int) error {
	_, errFree := GetAirdropFreeByID(id)
	if errFree == nil {
		return UpdateAirdropFreeByID(id, name, task, link, level, status, backed, funds, supply, fdv, marketCap, vesting, linkClaim, linkDiscord, linkTwitter, imageURL, description, linkGuide, price, usdIncome)
	}

	_, errPaid := GetAirdropPaidByID(id)
	if errPaid == nil {
		return UpdateAirdropPaidByID(id, name, task, link, level, status, backed, funds, supply, fdv, marketCap, vesting, linkClaim, linkDiscord, linkTwitter, imageURL, description, linkGuide, price, usdIncome)
	}

	return fmt.Errorf("UpdateAllAirdropByID: airdrop not found in both collections")
}

func UpdateAirdropFreeByID(id primitive.ObjectID, name, task, link, level, status, backed, funds, supply, fdv, marketCap, vesting, linkClaim, linkDiscord, linkTwitter, imageURL, description, linkGuide string, price float64, usdIncome int) error {
	collection := "airdrop_free"
	filter := bson.M{"_id": id}

	currentAirdrop, err := GetAirdropFreeByID(id)
	if err != nil {
		return fmt.Errorf("UpdateAirdropFreeByID: failed to get current airdrop: %v", err)
	}

	updateFields := bson.M{
		"name":         name,
		"task":         task,
		"link":         link,
		"level":        level,
		"status":       status,
		"backed":       backed,
		"funds":        funds,
		"supply":       supply,
		"fdv":          fdv,
		"market_cap":   marketCap,
		"vesting":      vesting,
		"link_claim":   linkClaim,
		"link_discord": linkDiscord,
		"link_twitter": linkTwitter,
		"image_url":    imageURL,
		"description":  description,
		"link_guide":   linkGuide,
		"price":        price,
		"usd_income":   usdIncome,
	}

	if status == "ended" && currentAirdrop.Status != "ended" {
		now := time.Now()
		updateFields["ended_at"] = now
	}

	update := bson.M{
		"$set": updateFields,
	}

	result, err := config.Database.Collection(collection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("UpdateAirdropFreeByID: %v", err)
	}

	if result.ModifiedCount == 0 {
		return errors.New("no data has been changed with the specified ID")
	}

	return nil
}

func UpdateAirdropPaidByID(id primitive.ObjectID, name, task, link, level, status, backed, funds, supply, fdv, marketCap, vesting, linkClaim, linkDiscord, linkTwitter, imageURL, description, linkGuide string, price float64, usdIncome int) error {
	collection := "airdrop_paid"
	filter := bson.M{"_id": id}

	currentAirdrop, err := GetAirdropPaidByID(id)
	if err != nil {
		return fmt.Errorf("UpdateAirdropPaidByID: failed to get current airdrop: %v", err)
	}

	updateFields := bson.M{
		"name":         name,
		"task":         task,
		"link":         link,
		"level":        level,
		"status":       status,
		"backed":       backed,
		"funds":        funds,
		"supply":       supply,
		"fdv":          fdv,
		"market_cap":   marketCap,
		"vesting":      vesting,
		"link_claim":   linkClaim,
		"link_discord": linkDiscord,
		"link_twitter": linkTwitter,
		"image_url":    imageURL,
		"description":  description,
		"link_guide":   linkGuide,
		"price":        price,
		"usd_income":   usdIncome,
	}

	if status == "ended" && currentAirdrop.Status != "ended" {
		now := time.Now()
		updateFields["ended_at"] = now
	}

	update := bson.M{
		"$set": updateFields,
	}

	result, err := config.Database.Collection(collection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("UpdateAirdropPaidByID: %v", err)
	}

	if result.ModifiedCount == 0 {
		return errors.New("no data has been changed with the specified ID")
	}

	return nil
}

func DeleteAllAirdropByID(id primitive.ObjectID) error {
	var errFree, errPaid error

	errFree = DeleteAirdropFreeByID(id)
	if errFree != nil {
		errPaid = DeleteAirdropPaidByID(id)
		if errPaid != nil {
			return fmt.Errorf("DeleteAllAirdropByID: airdrop not found in both collections. Free error: %v, Paid error: %v", errFree, errPaid)
		}
	}

	return nil
}

func DeleteAirdropFreeByID(id primitive.ObjectID) error {
	collection := config.Database.Collection("airdrop_free")
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s in airdrop_free: %s", id.Hex(), err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found in airdrop_free", id.Hex())
	}

	return nil
}

func DeleteAirdropPaidByID(id primitive.ObjectID) error {
	collection := config.Database.Collection("airdrop_paid")
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s in airdrop_paid: %s", id.Hex(), err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found in airdrop_paid", id.Hex())
	}

	return nil
}
