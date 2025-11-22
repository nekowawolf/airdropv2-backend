package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/nekowawolf/airdropv2/models"
	"github.com/nekowawolf/airdropv2/module"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/nekowawolf/airdropv2/config"
)

var imageCollection = config.Database.Collection("images")

func UploadImageHandler(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "file is required"})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	url, sha, path, err := module.UploadToGitHub(file, fileHeader)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	img := models.Image{
		ID:       primitive.NewObjectID(),
		Filename: fileHeader.Filename,
		URL:      url,
		Size:     fileHeader.Size,
		Sha:      sha,
		Path:     path,
	}

	_, err = imageCollection.InsertOne(context.Background(), img)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to save db"})
	}

	return c.JSON(img)
}

func GetAllImages(c *fiber.Ctx) error {
	cursor, err := imageCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch"})
	}

	var results []models.Image
	cursor.All(context.Background(), &results)

	return c.JSON(results)
}

func DeleteImage(c *fiber.Ctx) error {
	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	var img models.Image
	err = imageCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&img)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}

	err = module.DeleteFromGitHub(img.Path, img.Sha)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	_, err = imageCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed delete db"})
	}

	return c.JSON(fiber.Map{"message": "deleted"})
}
