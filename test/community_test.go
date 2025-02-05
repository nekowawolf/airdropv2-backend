package test

import (
	"fmt"
	"testing"

	"github.com/nekowawolf/airdropv2/models"
	"github.com/nekowawolf/airdropv2/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertCryptoCommunity(t *testing.T) {
	name := "Nusan Airdrop"
	platforms := "Telegram"
	category := "Airdrop"
	imgURL := "https://example.com/image.png"
	linkURL := "https://example.com"

	result := module.InsertCryptoCommunity(name, platforms, category, imgURL, linkURL)

	id, ok := result.(primitive.ObjectID)
	if !ok || id.IsZero() {
		t.Errorf("Failed to insert crypto community: invalid ID returned")
		return
	}

	fmt.Printf("Inserted CryptoCommunity ID: %v\n", id.Hex())
}

func TestGetAllCryptoCommunity(t *testing.T) {
	data, err := module.GetAllCryptoCommunity()
	if err != nil {
		t.Errorf("Failed to retrieve crypto communities: %v", err)
	} else if len(data) == 0 {
		t.Errorf("No crypto communities found")
	} else {
		fmt.Printf("Retrieved crypto communities: %v\n", data)
	}
}

func TestGetCryptoCommunityByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("675c07d48350607d1c13534f") 
	if err != nil {
		t.Fatalf("Invalid ObjectID: %v", err)
	}

	notes, err := module.GetCryptoCommunityByID(id)
	if err != nil {
		t.Errorf("Failed to get notes: %v", err)
	} else {
		t.Logf("Retrieved Notes: %+v", notes)
	}
}

func TestGetCryptoCommunityByName(t *testing.T) {
	name := "Airdrop"

	communities, err := module.GetCryptoCommunityByName(name)
	if err != nil {
		t.Fatalf("Error calling GetCryptoCommunityByName: %v", err)
	}

	fmt.Println("Crypto Communities found:", communities)
}

func TestUpdateCryptoCommunityByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("675bf50331909619aeb35400")
	if err != nil {
		t.Fatalf("Invalid ObjectID: %v", err)
	}

	newName := "N Airdrop"
	newPlatforms := "Discord"
	newCategory := "NFT"
	newImgURL := "https://example.com/updated-image.png"
	newLinkURL := "https://example.com/updated-link"

	updatedCrypto, err := module.UpdateCryptoCommunityByID(id, models.CryptoCommunity{
		Name:      newName,
		Platforms: newPlatforms,
		Category:  newCategory,
		ImgURL:    newImgURL,
		LinkURL:   newLinkURL,
	})
	if err != nil {
		t.Errorf("Failed to update crypto community: %v", err)
		return
	}

	t.Logf("Successfully updated crypto community: %+v", updatedCrypto)
}

func TestDeleteCryptoCommunityByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("675bf50331909619aeb35400")
	if err != nil {
		t.Fatalf("Invalid ObjectID: %v", err)
	}

	err = module.DeleteCryptoCommunityByID(id)
	if err != nil {
		t.Errorf("Failed to delete CryptoCommunity by ID: %v", err)
		return
	}

	t.Logf("CryptoCommunity with ID %s deleted successfully", id.Hex())
}