package portfolio

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/nekowawolf/airdropv2/utils"
)

func GetPortfolioHandler(c *fiber.Ctx) error {
	data, err := utils.GetOrSetCache("portfolio", 24*time.Hour, func() (*Portfolio, error) {
		return GetPortfolio()
	})
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Portfolio not found"})
	}
	return c.JSON(data)
}

func UpdatePortfolioHandler(c *fiber.Ctx) error {
	var req Portfolio
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}
	if err := UpdatePortfolio(req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Update failed"})
	}
	utils.InvalidateCache("portfolio")
	return c.JSON(fiber.Map{"message": "Portfolio updated"})
}

func UpdateHeroProfileHandler(c *fiber.Ctx) error {
	var req HeroProfile
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}
	if err := UpdateHeroProfile(req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Update hero profile failed"})
	}
	utils.InvalidateCache("portfolio")
	return c.JSON(fiber.Map{"message": "Hero profile updated"})
}

func AddExperienceHandler(c *fiber.Ctx) error {
	var req Experience
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}
	if err := AddExperience(req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to add experience"})
	}
	utils.InvalidateCache("portfolio")
	return c.JSON(fiber.Map{"message": "Experience added"})
}

func AddEducationHandler(c *fiber.Ctx) error {
	var req Education
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}
	if err := AddEducation(req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to add education"})
	}
	utils.InvalidateCache("portfolio")
	return c.JSON(fiber.Map{"message": "Education added"})
}

func AddTechSkillHandler(c *fiber.Ctx) error {
	var req SkillItem
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}
	if err := AddTechSkill(req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to add tech skill"})
	}
	utils.InvalidateCache("portfolio")
	return c.JSON(fiber.Map{"message": "Tech skill added"})
}

func DeleteExperienceHandler(c *fiber.Ctx) error {
	if err := DeleteExperience(c.Params("id")); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete experience"})
	}
	utils.InvalidateCache("portfolio")
	return c.JSON(fiber.Map{"message": "Experience deleted"})
}

func DeleteEducationHandler(c *fiber.Ctx) error {
	if err := DeleteEducation(c.Params("id")); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete education"})
	}
	utils.InvalidateCache("portfolio")
	return c.JSON(fiber.Map{"message": "Education deleted"})
}

func AddDesignSkillHandler(c *fiber.Ctx) error {
	var req SkillItem
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}
	if err := AddDesignSkill(req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to add design skill"})
	}
	utils.InvalidateCache("portfolio")
	return c.JSON(fiber.Map{"message": "Design skill added"})
}

func DeleteDesignSkillHandler(c *fiber.Ctx) error {
	if err := DeleteDesignSkill(c.Params("id")); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete design skill"})
	}
	utils.InvalidateCache("portfolio")
	return c.JSON(fiber.Map{"message": "Design skill deleted"})
}

func DeleteTechSkillHandler(c *fiber.Ctx) error {
	if err := DeleteTechSkill(c.Params("id")); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete tech skill"})
	}
	utils.InvalidateCache("portfolio")
	return c.JSON(fiber.Map{"message": "Tech skill deleted"})
}