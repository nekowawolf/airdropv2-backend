package portfolio

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/nekowawolf/airdropv2/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func invalidateProjectCache() {
	utils.InvalidateCache("portfolio_projects")
}

func GetProjectsHandler(c *fiber.Ctx) error {
	data, err := utils.GetOrSetCache("portfolio_projects", 24*time.Hour, func() ([]Project, error) {
		return GetAllProjects()
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve projects",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    data,
	})
}

func GetProjectByIDHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	data, err := GetProjectByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Project not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    data,
	})
}

func InsertProjectHandler(c *fiber.Ctx) error {
	var req Project
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	insertedID, err := InsertProject(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert project",
		})
	}

	if objectID, ok := insertedID.(primitive.ObjectID); ok {
		invalidateProjectCache()
		return c.JSON(fiber.Map{
			"message":     "Project inserted successfully",
			"inserted_id": objectID.Hex(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Failed to retrieve inserted ID",
	})
}

func UpdateProjectHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	var req Project
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	if err := UpdateProjectByID(id, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update project",
		})
	}

	invalidateProjectCache()
	return c.JSON(fiber.Map{"message": "Project updated successfully"})
}

func DeleteProjectHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	if err := DeleteProjectByID(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete project",
		})
	}

	invalidateProjectCache()
	return c.JSON(fiber.Map{"message": "Project deleted successfully"})
}