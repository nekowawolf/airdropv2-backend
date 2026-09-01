package portfolio

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/nekowawolf/airdropv2/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func invalidateDesignCache() {
	utils.InvalidateCache("portfolio_designs")
}

func GetDesignsHandler(c *fiber.Ctx) error {
	data, err := utils.GetOrSetCache("portfolio_designs", 24*time.Hour, func() ([]Design, error) {
		return GetAllDesigns()
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve designs",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    data,
	})
}

func GetDesignByIDHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	data, err := GetDesignByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Design not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    data,
	})
}

func InsertDesignHandler(c *fiber.Ctx) error {
	var req Design
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	insertedID, err := InsertDesign(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert design",
		})
	}

	if objectID, ok := insertedID.(primitive.ObjectID); ok {
		invalidateDesignCache()
		return c.JSON(fiber.Map{
			"message":     "Design inserted successfully",
			"inserted_id": objectID.Hex(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Failed to retrieve inserted ID",
	})
}

func UpdateDesignHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	var req Design
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	if err := UpdateDesignByID(id, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update design",
		})
	}

	invalidateDesignCache()
	return c.JSON(fiber.Map{"message": "Design updated successfully"})
}

func DeleteDesignHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	if err := DeleteDesignByID(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete design",
		})
	}

	invalidateDesignCache()
	return c.JSON(fiber.Map{"message": "Design deleted successfully"})
}