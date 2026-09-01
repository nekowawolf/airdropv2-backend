package notes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/nekowawolf/airdropv2/utils"
)

func GetAllNotesHandler(c *fiber.Ctx) error {
	notes, err := GetAllNotes()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    notes,
	})
}

func GetNoteByIDHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	note, err := GetNoteByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Note not found",
		})
	}

	return c.JSON(note)
}

func InsertNoteHandler(c *fiber.Ctx) error {
	var req Notes

	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	insertedID := InsertNote(
		req.Title,
		req.Content,
		req.Type,
	)

	if insertedID == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert Note",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "Note created successfully",
		"insertedID": insertedID,
	})
}

func UpdateNoteByIDHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	var req Notes

	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	updateData := Notes{
		Title:   req.Title,
		Content: req.Content,
		Type:    req.Type,
	}

	updatedNote, err := UpdateNoteByID(id, updateData)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Note not found or could not be updated",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Note updated successfully",
		"data":    updatedNote,
	})
}

func DeleteNoteByIDHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	err = DeleteNoteByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Note deleted successfully",
	})
}