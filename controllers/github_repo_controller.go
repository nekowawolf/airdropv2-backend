package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nekowawolf/airdropv2/models"
	"github.com/nekowawolf/airdropv2/module"
	"github.com/nekowawolf/airdropv2/utils"
)

func invalidateGithubRepoCache() {
	utils.InvalidateCache("githubrepo", "githubrepo_stats")
}

func GetAllGithubRepos(c *fiber.Ctx) error {
	repos, err := utils.GetOrSetCache("githubrepo", 24*time.Hour, func() ([]models.GithubRepo, error) {
		return module.GetAllGithubRepos()
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    repos,
	})
}

func GetGithubRepoStats(c *fiber.Ctx) error {
	stats, err := utils.GetOrSetCache("githubrepo_stats", 24*time.Hour, func() (map[string]interface{}, error) {
		return module.GetGithubRepoStats()
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

func GetGithubRepoByID(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	repo, err := module.GetGithubRepoByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "GithubRepo not found",
		})
	}

	return c.JSON(repo)
}

func InsertGithubRepo(c *fiber.Ctx) error {
	var req models.GithubRepo

	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	insertedID := module.InsertGithubRepo(
		req.Name,
		req.Description,
		req.Category,
		req.RepoURL,
		req.Owner,
		req.RepoName,
		req.Website,
		req.Twitter,
		req.Instagram,
		req.Discord,
		req.AddedBy,
	)

	if insertedID == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert GithubRepo",
		})
	}

	invalidateGithubRepoCache()
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "GithubRepo created successfully",
		"insertedID": insertedID,
	})
}

func UpdateGithubRepoByID(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	var req models.GithubRepo

	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	updateData := models.GithubRepo{
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		RepoURL:     req.RepoURL,
		Owner:       req.Owner,
		RepoName:    req.RepoName,
		Website:     req.Website,
		Twitter:     req.Twitter,
		Instagram:   req.Instagram,
		Discord:     req.Discord,
		AddedBy:     req.AddedBy,
	}

	updatedRepo, err := module.UpdateGithubRepoByID(id, updateData)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "GithubRepo not found or could not be updated",
		})
	}

	invalidateGithubRepoCache()
	return c.JSON(fiber.Map{
		"message": "GithubRepo updated successfully",
		"data":    updatedRepo,
	})
}

func DeleteGithubRepoByID(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	err = module.DeleteGithubRepoByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	invalidateGithubRepoCache()
	return c.JSON(fiber.Map{
		"message": "GithubRepo deleted successfully",
	})
}

func GetGithubRepoHistory(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	period := c.Query("period", "month")
	var targetTime time.Time
	now := time.Now().UTC()

	switch period {
	case "today":
		targetTime = now.Add(-24 * time.Hour)
	case "week":
		targetTime = now.Add(-7 * 24 * time.Hour)
	case "month":
		targetTime = now.Add(-30 * 24 * time.Hour)
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid period. Must be 'today', 'week', or 'month'",
		})
	}

	cacheKey := "githubrepo_history_" + id.Hex() + "_" + period

	type HistoryResponse struct {
		Period    string `json:"period"`
		Available bool   `json:"available"`
		Current   any    `json:"current"`
		Previous  any    `json:"previous"`
		Growth    any    `json:"growth"`
	}

	responseData, err := utils.GetOrSetCache(cacheKey, 24*time.Hour, func() (HistoryResponse, error) {
		repo, err := module.GetGithubRepoByID(id)
		if err != nil {
			return HistoryResponse{}, err
		}

		currentStats := map[string]int{
			"stars": 0,
			"forks": 0,
		}
		if repo.Stats != nil {
			currentStats["stars"] = repo.Stats.Stars
			currentStats["forks"] = repo.Stats.Forks
		}

		history, _ := module.GetGithubRepoHistoryByTargetTime(id, targetTime)
		
		if history == nil {
			return HistoryResponse{
				Period:    period,
				Available: false,
				Current:   currentStats,
				Previous:  nil,
				Growth:    nil,
			}, nil
		}

		previousStats := map[string]int{
			"stars": history.Stars,
			"forks": history.Forks,
		}

		growthStats := map[string]int{
			"stars": currentStats["stars"] - history.Stars,
			"forks": currentStats["forks"] - history.Forks,
		}

		return HistoryResponse{
			Period:    period,
			Available: true,
			Current:   currentStats,
			Previous:  previousStats,
			Growth:    growthStats,
		}, nil
	})

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "GithubRepo not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "History retrieved successfully",
		"data":    responseData,
	})
}

func GetGithubRepoDetails(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	cacheKey := "githubrepo_details_" + id.Hex()

	responseData, err := utils.GetOrSetCache(cacheKey, 24*time.Hour, func() (models.GithubRepoDetailsData, error) {
		repo, err := module.GetGithubRepoByID(id)
		if err != nil {
			return models.GithubRepoDetailsData{}, err
		}

		repoData, mdFiles, err := utils.FetchGithubRepoDetails(repo.Owner, repo.RepoName)
		if err != nil {
			return models.GithubRepoDetailsData{}, err
		}

		return models.GithubRepoDetailsData{
			RepoData: repoData,
			MdFiles:  mdFiles,
		}, nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch Github Repo details: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Details retrieved successfully",
		"data":    responseData,
	})
}

func GetGithubRepoCommits(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	perPage := c.Query("per_page", "8")
	cacheKey := "githubrepo_commits_" + id.Hex() + "_" + perPage

	responseData, err := utils.GetOrSetCache(cacheKey, 1*time.Hour, func() ([]interface{}, error) {
		repo, err := module.GetGithubRepoByID(id)
		if err != nil {
			return nil, err
		}

		commits, err := utils.FetchGithubCommits(repo.Owner, repo.RepoName, perPage)
		if err != nil {
			return nil, err
		}

		return commits, nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch Github Repo commits: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Commits retrieved successfully",
		"data":    responseData,
	})
}

func GetGithubRepoCommitsByOwnerRepo(c *fiber.Ctx) error {
	owner := c.Params("owner")
	repoName := c.Params("repoName")
	
	if owner == "" || repoName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Owner and repoName are required",
		})
	}

	perPage := c.Query("per_page", "8")
	cacheKey := "githubrepo_commits_raw_" + owner + "_" + repoName + "_" + perPage

	responseData, err := utils.GetOrSetCache(cacheKey, 1*time.Hour, func() ([]interface{}, error) {
		commits, err := utils.FetchGithubCommits(owner, repoName, perPage)
		if err != nil {
			return nil, err
		}
		return commits, nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch Github Repo commits: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Commits retrieved successfully",
		"data":    responseData,
	})
}