package controllers

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nekowawolf/airdropv2/config"
	"github.com/nekowawolf/airdropv2/models"
	"github.com/nekowawolf/airdropv2/module"
	"github.com/nekowawolf/airdropv2/utils"
)



func SubmitRepo(c *fiber.Ctx) error {
	ip := c.IP()
	key := fmt.Sprintf("rate:repo_submission:%s", ip)
	ctx := context.Background()

	count, err := config.RedisClient.Incr(ctx, key).Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error during rate limiting",
		})
	}

	if count == 1 {
		config.RedisClient.Expire(ctx, key, 5*time.Minute)
	}

	if count > 10 {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"error": "Too many requests",
		})
	}

	var req models.RepoSubmissionRequest
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	if req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name is required",
		})
	}

	if !strings.Contains(req.RepoURL, "github.com") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid GitHub repository URL",
		})
	}

	if _, err := url.ParseRequestURI(req.RepoURL); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL format for RepoURL",
		})
	}

	if req.Link != "" {
		if _, err := url.ParseRequestURI(req.Link); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid URL format for Link",
			})
		}
	}

	if !utils.VerifyTurnstile(req.TurnstileToken) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Invalid or missing Turnstile token",
		})
	}

	addedBy := &models.AddedByInfo{
		Name: req.Name,
		URL:  req.Link,
	}

	insertedID := module.InsertRepoSubmission(req.RepoURL, addedBy)
	if insertedID == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to submit repository",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "Repository submitted successfully",
		"insertedID": insertedID,
	})
}

func GetAllRepoSubmissions(c *fiber.Ctx) error {
	submissions, err := module.GetAllRepoSubmissions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    submissions,
	})
}

func DeleteRepoSubmission(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	err = module.DeleteRepoSubmissionByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Repo submission deleted successfully",
	})
}