package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nekowawolf/airdropv2/module"
	"github.com/nekowawolf/airdropv2/models"
)

func GetPortfolio(c *fiber.Ctx) error {
	data, err := module.GetPortfolio()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Portfolio not found"})
	}
	return c.JSON(data)
}

func UpdatePortfolio(c *fiber.Ctx) error {
	var req models.Portfolio
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}
	if err := module.UpdatePortfolio(req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Update failed"})
	}
	return c.JSON(fiber.Map{"message": "Portfolio updated"})
}

func UpdateHeroProfile(c *fiber.Ctx) error {
	var req models.HeroProfile
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}
	if err := module.UpdateHeroProfile(req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Update hero profile failed"})
	}
	return c.JSON(fiber.Map{"message": "Hero profile updated"})
}

func AddCertificate(c *fiber.Ctx) error {
	var req models.Certificate
	c.BodyParser(&req)
	module.AddCertificate(req)
	return c.JSON(fiber.Map{"message": "Certificate added"})
}

func AddDesign(c *fiber.Ctx) error {
	var req models.Design
	c.BodyParser(&req)
	module.AddDesign(req)
	return c.JSON(fiber.Map{"message": "Design added"})
}

func AddProject(c *fiber.Ctx) error {
	var req models.Project
	c.BodyParser(&req)
	module.AddProject(req)
	return c.JSON(fiber.Map{"message": "Project added"})
}

func AddExperience(c *fiber.Ctx) error {
	var req models.Experience
	c.BodyParser(&req)
	module.AddExperience(req)
	return c.JSON(fiber.Map{"message": "Experience added"})
}

func AddEducation(c *fiber.Ctx) error {
	var req models.Education
	c.BodyParser(&req)
	module.AddEducation(req)
	return c.JSON(fiber.Map{"message": "Education added"})
}

func AddTechSkill(c *fiber.Ctx) error {
	var req models.SkillItem
	c.BodyParser(&req)
	module.AddTechSkill(req)
	return c.JSON(fiber.Map{"message": "Tech skill added"})
}

func DeleteCertificate(c *fiber.Ctx) error {
	module.DeleteCertificate(c.Params("id"))
	return c.JSON(fiber.Map{"message": "Certificate deleted"})
}

func DeleteDesign(c *fiber.Ctx) error {
	module.DeleteDesign(c.Params("id"))
	return c.JSON(fiber.Map{"message": "Design deleted"})
}

func DeleteProject(c *fiber.Ctx) error {
	module.DeleteProject(c.Params("id"))
	return c.JSON(fiber.Map{"message": "Project deleted"})
}

func DeleteExperience(c *fiber.Ctx) error {
	module.DeleteExperience(c.Params("id"))
	return c.JSON(fiber.Map{"message": "Experience deleted"})
}

func DeleteEducation(c *fiber.Ctx) error {
	module.DeleteEducation(c.Params("id"))
	return c.JSON(fiber.Map{"message": "Education deleted"})
}

func DeleteTechSkill(c *fiber.Ctx) error {
	module.DeleteTechSkill(c.Params("id"))
	return c.JSON(fiber.Map{"message": "Tech skill deleted"})
}