package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nekowawolf/airdropv2/models"
	"github.com/nekowawolf/airdropv2/module"
	"github.com/nekowawolf/airdropv2/utils"
)

func invalidateNetCache() {
	utils.InvalidateCache("net", "net_stats")
}

func GetAllNet(c *fiber.Ctx) error {
	nets, err := utils.GetOrSetCache("net", 24*time.Hour, func() ([]models.Net, error) {
		return module.GetAllNet()
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    nets,
	})
}

func GetNetStats(c *fiber.Ctx) error {
	stats, err := utils.GetOrSetCache("net_stats", 24*time.Hour, func() (map[string]interface{}, error) {
		return module.GetNetStats()
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Stats retrieved successfully",
		"data":    stats,
	})
}

func GetNetByID(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	netItem, err := module.GetNetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Net not found",
		})
	}

	return c.JSON(netItem)
}

func InsertNet(c *fiber.Ctx) error {
	var req models.Net

	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	insertedID := module.InsertNet(
		req.Name,
		req.Description,
		req.ImageURL,
		req.Website,
		req.Categories,
		req.Media,
		req.Socials,
	)

	if insertedID == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert Net",
		})
	}

	invalidateNetCache()
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "Net created successfully",
		"insertedID": insertedID,
	})
}

func UpdateNetByID(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	var req models.Net

	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	updateData := models.Net{
		Name:           req.Name,
		Description:    req.Description,
		ImageURL:       req.ImageURL,
		Website:        req.Website,
		Categories:     req.Categories,
		Media:          req.Media,
		Socials:        req.Socials,
	}

	updatedNet, err := module.UpdateNetByID(id, updateData)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Net not found or could not be updated",
		})
	}

	invalidateNetCache()
	return c.JSON(fiber.Map{
		"message": "Net updated successfully",
		"data":    updatedNet,
	})
}

func DeleteNetByID(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	err = module.DeleteNetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	invalidateNetCache()
	return c.JSON(fiber.Map{
		"message": "Net deleted successfully",
	})
}