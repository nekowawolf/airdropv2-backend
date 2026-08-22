package controllers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/nekowawolf/airdropv2/module"
	"github.com/nekowawolf/airdropv2/utils"
)

func UploadMediaHandler(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "file is required",
		})
	}

	contentType := fileHeader.Header.Get("Content-Type")
	mediaType := utils.GetMediaType(contentType)

	if mediaType == "unknown" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "unsupported file type, only images and videos are allowed",
		})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to open file",
		})
	}
	defer file.Close()

	fileBytes, err := utils.StreamToBytes(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to read file",
		})
	}

	filename := fileHeader.Filename
	uploadBytes := fileBytes
	uploadContentType := contentType

	if utils.IsImageContentType(contentType) && !strings.HasSuffix(strings.ToLower(filename), ".webp") {
		webpBytes, err := utils.ConvertToWebP(fileBytes)
		if err != nil {
			fmt.Printf("WebP conversion failed, uploading original: %v\n", err)
		} else {
			uploadBytes = webpBytes
			uploadContentType = "image/webp"
			if dotIdx := strings.LastIndex(filename, "."); dotIdx != -1 {
				filename = filename[:dotIdx] + ".webp"
			} else {
				filename = filename + ".webp"
			}
		}
	}

	r2Key := utils.GenerateR2Key(mediaType, filename)

	url, err := utils.UploadToR2(uploadBytes, r2Key, uploadContentType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	insertedID := module.InsertMedia(
		filename,
		url,
		int64(len(uploadBytes)),
		uploadContentType,
		mediaType,
		r2Key,
	)

	if insertedID == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save media to database",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Media uploaded successfully",
		"data": fiber.Map{
			"_id":          insertedID,
			"filename":     filename,
			"url":          url,
			"size":         int64(len(uploadBytes)),
			"content_type": uploadContentType,
			"media_type":   mediaType,
			"r2_key":       r2Key,
		},
	})
}

func GetAllMedia(c *fiber.Ctx) error {
	mediaList, err := module.GetAllMedia()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    mediaList,
	})
}

func DeleteMedia(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	media, err := module.GetMediaByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Media not found",
		})
	}

	err = utils.DeleteFromR2(media.R2Key)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = module.DeleteMediaByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Media deleted successfully",
	})
}