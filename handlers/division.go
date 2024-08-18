package handlers

import (
	"errors"
	"info-bd-api/database"
	"info-bd-api/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllDivision(c *fiber.Ctx) error {

	var div []model.Division

	res := database.DB.Find(&div)

	if res.Error != nil {
		return c.JSON(fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch divisions"))
	}

	return c.JSON(div)
}

func GetDivisionById(c *fiber.Ctx) error {
	id := c.Params("id")

	var division model.Division

	if err := database.DB.First(&division, id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Division not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch division",
		})
	}

	return c.Status(fiber.StatusOK).JSON(division)
}

func GetDivisionByName(c *fiber.Ctx) error {

	name := c.Params("name")
	var div model.Division
	if err := database.DB.Where("name = ?", name).First(&div).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Division not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch division",
		})
	}
	return c.Status(fiber.StatusOK).JSON(div)
}

func AddDivision(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "district added"))
}

func UpdateDivisionById(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "Update Division By id"))
}
func UpdateDivisionByName(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "Update Division By Name"))
}
func DeleteDivisionByName(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "Delete Division By Name"))
}
func DeleteDivisionById(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "Delete Division By id"))
}
