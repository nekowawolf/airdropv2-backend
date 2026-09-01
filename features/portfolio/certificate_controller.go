package portfolio

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/nekowawolf/airdropv2/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func invalidateCertificateCache() {
	utils.InvalidateCache("portfolio_certificates")
}

func GetCertificatesHandler(c *fiber.Ctx) error {
	data, err := utils.GetOrSetCache("portfolio_certificates", 24*time.Hour, func() ([]Certificate, error) {
		return GetAllCertificates()
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve certificates",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    data,
	})
}

func GetCertificateByIDHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	data, err := GetCertificateByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Certificate not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    data,
	})
}

func InsertCertificateHandler(c *fiber.Ctx) error {
	var req Certificate
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	insertedID, err := InsertCertificate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert certificate",
		})
	}

	if objectID, ok := insertedID.(primitive.ObjectID); ok {
		invalidateCertificateCache()
		return c.JSON(fiber.Map{
			"message":     "Certificate inserted successfully",
			"inserted_id": objectID.Hex(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Failed to retrieve inserted ID",
	})
}

func UpdateCertificateHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	var req Certificate
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	if err := UpdateCertificateByID(id, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update certificate",
		})
	}

	invalidateCertificateCache()
	return c.JSON(fiber.Map{"message": "Certificate updated successfully"})
}

func DeleteCertificateHandler(c *fiber.Ctx) error {
	id, err := utils.ParseObjectID(c, "id")
	if err != nil {
		return err
	}

	if err := DeleteCertificateByID(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete certificate",
		})
	}

	invalidateCertificateCache()
	return c.JSON(fiber.Map{"message": "Certificate deleted successfully"})
}