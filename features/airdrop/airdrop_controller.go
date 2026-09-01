package airdrop

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/nekowawolf/airdropv2/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func invalidateAirdropCache() {
	utils.InvalidateCache("airdrops_all", "airdrops_free", "airdrops_paid", "airdrops_stats")
}

// PUBLIC ENDPOINTS
func GetAirdropsPublicHandler(c *fiber.Ctx) error {
	isPaidStr := c.Query("is_paid")
	
	cacheKey := "airdrops_all"
	var isPaidFilter *bool
	
	if isPaidStr == "true" {
		cacheKey = "airdrops_paid"
		val := true
		isPaidFilter = &val
	} else if isPaidStr == "false" {
		cacheKey = "airdrops_free"
		val := false
		isPaidFilter = &val
	}

	data, err := utils.GetOrSetCache(cacheKey, 24*time.Hour, func() ([]AirdropPublic, error) {
		adminData, err := GetAirdrops(isPaidFilter)
		if err != nil {
			return nil, err
		}
		
		publicData := make([]AirdropPublic, len(adminData))
		for i, a := range adminData {
			publicData[i] = a.AirdropPublic
		}
		
		if publicData == nil {
			publicData = []AirdropPublic{}
		}
		
		return publicData, nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve Airdrops data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    data,
	})
}

func GetAirdropsStatsHandler(c *fiber.Ctx) error {
	stats, err := utils.GetOrSetCache("airdrops_stats", 24*time.Hour, func() (map[string]int, error) {
		return GetAirdropStats()
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Stats retrieved successfully",
		"data":    stats,
	})
}

// ADMIN ENDPOINTS 
func GetAirdropsAdminHandler(c *fiber.Ctx) error {
	isPaidStr := c.Query("is_paid")
	var isPaidFilter *bool
	
	if isPaidStr == "true" {
		val := true
		isPaidFilter = &val
	} else if isPaidStr == "false" {
		val := false
		isPaidFilter = &val
	}

	data, err := GetAirdrops(isPaidFilter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve Airdrops data",
		})
	}

	if data == nil {
		data = []AirdropAdmin{}
	}

	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    data,
	})
}

func GetAirdropByIDAdminHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	data, err := GetAirdropByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Airdrop not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    data,
	})
}

func InsertAirdropHandler(c *fiber.Ctx) error {
	var reqAirdrop AirdropAdmin

	if err := utils.ParseBody(c, &reqAirdrop); err != nil {
		return err
	}

	if reqAirdrop.Status == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Status is required",
		})
	}

	insertedID, err := InsertAirdrop(reqAirdrop)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert Airdrop",
		})
	}

	if objectID, ok := insertedID.(primitive.ObjectID); ok {
		invalidateAirdropCache()
		return c.JSON(fiber.Map{
			"message":     "Airdrop inserted successfully",
			"inserted_id": objectID.Hex(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Failed to retrieve inserted ID",
	})
}

func UpdateAirdropHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	var updateData AirdropAdmin

	if err := utils.ParseBody(c, &updateData); err != nil {
		return err
	}

	err = UpdateAirdropByID(id, updateData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update Airdrop by ID: " + err.Error(),
		})
	}

	invalidateAirdropCache()
	return c.JSON(fiber.Map{
		"message": "Airdrop updated successfully",
	})
}

func DeleteAirdropHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	err = DeleteAirdropByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	invalidateAirdropCache()
	return c.JSON(fiber.Map{
		"message": "Airdrop deleted successfully",
	})
}