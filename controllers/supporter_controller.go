package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nekowawolf/airdropv2/models"
	"github.com/nekowawolf/airdropv2/module"
	"github.com/nekowawolf/airdropv2/utils"
)

func GetAllSupporters(c *fiber.Ctx) error {
	supporters, err := module.GetAllSupporters()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    supporters,
	})
}

func GetSupporterByID(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	supporter, err := module.GetSupporterByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    supporter,
	})
}

func InsertSupporter(c *fiber.Ctx) error {
	type SupporterRequest struct {
		models.Supporter
		TurnstileToken string `json:"turnstileToken"`
	}

	var req SupporterRequest
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	if !utils.VerifyTurnstile(req.TurnstileToken) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Invalid or missing Turnstile token",
		})
	}

	insertedID := module.InsertSupporter(&req.Supporter)
	if insertedID == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert supporter",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "Supporter created successfully",
		"insertedID": insertedID,
	})
}

func UpdateSupporterByID(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	var req models.Supporter
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	err = module.UpdateSupporterByID(id, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Supporter updated successfully",
	})
}

func DeleteSupporterByID(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	err = module.DeleteSupporterByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Supporter deleted successfully",
	})
}