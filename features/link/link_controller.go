package link

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/nekowawolf/airdropv2/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ==================== PROFILE CONTROLLERS ====================

func GetProfileHandler(c *fiber.Ctx) error {
	profile, err := utils.GetOrSetCache("profilelink", 24*time.Hour, func() (*Profile, error) {
		return GetProfile()
	})
	if err != nil {
		// Return default profile if not found
		return c.JSON(fiber.Map{
			"name":       "nekowawolf",
			"username":   "nekowawolf",
			"bio":        "Professional Coder (vibe coding)",
			"avatar_url": "https://nekowawolf.github.io/cdn-images/images/2025/1763530019_113094795.jpeg",
			"cover_url":  "https://nekowawolf.github.io/cdn-images/images/2026/1775599464_bg_link.png",
			"links": fiber.Map{
				"github":    "https://github.com/nekowawolf",
				"twitter":   "https://x.com/nekowawolf_",
				"tiktok":    "https://tiktok.com/@nekowawolf",
				"website":   "https://nekowawolf.xyz/",
				"instagram": "https://instagram.com/nekowawolf",
			},
		})
	}

	return c.JSON(fiber.Map{
		"name":       profile.Name,
		"username":   profile.Username,
		"bio":        profile.Bio,
		"avatar_url": profile.AvatarURL,
		"cover_url":  profile.CoverURL,
		"links":      profile.Links,
	})
}

func UpdateProfileHandler(c *fiber.Ctx) error {
	var req Profile
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	if err := UpdateProfile(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	utils.InvalidateCache("profilelink")
	return c.JSON(fiber.Map{
		"message": "Profile updated successfully",
	})
}

func GetPostStatsHandler(c *fiber.Ctx) error {
	stats, err := utils.GetOrSetCache("postslink_stats", 24*time.Hour, func() (interface{}, error) {
		return GetPostStats()
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve stats",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Stats retrieved successfully",
		"data":    stats,
	})
}

// ==================== POSTS CONTROLLERS ====================

func GetAllPostsHandler(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 5)
	category := c.Query("category", "")
	search := c.Query("search", "")

	cacheKey := fmt.Sprintf("postslink:%d:%d:%s:%s", page, limit, category, search)

	posts, err := utils.GetOrSetCache(cacheKey, 24*time.Hour, func() (interface{}, error) {
		return GetPostsPaginated(page, limit, category, search)
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve posts",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Posts retrieved successfully",
		"data":    posts,
	})
}

func GetPostByIDHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	post, err := GetPostByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Post not found",
		})
	}

	sessionID := c.Get("X-Session-ID")
	if sessionID == "" {
		sessionID = c.IP() + c.Get("User-Agent")
	}

	_ = IncrementPostView(id, sessionID)

	viewCount, _ := GetPostViewCount(id)
	post.Views = int(viewCount)

	return c.JSON(fiber.Map{
		"message": "Post retrieved successfully",
		"data":    post,
	})
}

func CreatePostHandler(c *fiber.Ctx) error {
	var req LinkPost
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	if req.Caption == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Caption is required",
		})
	}

	if req.Category == "" {
		req.Category = "all"
	}

	insertedID, err := InsertPost(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if objectID, ok := insertedID.(primitive.ObjectID); ok {
		utils.InvalidateCache("postslink_stats")
		utils.InvalidateCachePrefix("postslink:")
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message":     "Post created successfully",
			"inserted_id": objectID.Hex(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Failed to retrieve inserted ID",
	})
}

func UpdatePostHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	var req LinkPost
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	if err := UpdatePost(id, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	utils.InvalidateCache("postslink_stats")
	utils.InvalidateCachePrefix("postslink:")
	return c.JSON(fiber.Map{
		"message": "Post updated successfully",
	})
}

func DeletePostHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	if err := DeletePost(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	utils.InvalidateCache("postslink_stats")
	utils.InvalidateCachePrefix("postslink:")
	return c.JSON(fiber.Map{
		"message": "Post deleted successfully",
	})
}