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

	var division model.Division
	if err := c.BodyParser(&division); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse request",
		})
	}
	if err := database.DB.Create(&division).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to save data",
		})
	}

	return c.JSON(fiber.NewError(fiber.StatusCreated, "division added"))
}

func UpdateDivisionById(c *fiber.Ctx) error {
	id := c.Params("id")
	var division model.Division

	if err := database.DB.First(&division, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Division not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve division",
		})
	}

	if err := c.BodyParser(&division); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	if err := database.DB.Save(&division).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update division",
		})
	}

	return c.JSON(fiber.NewError(fiber.StatusOK, "division updated"))
}

func UpdateDivisionByName(c *fiber.Ctx) error {
	name := c.Params("name")
	var division model.District

	if err := database.DB.Where("name = ?", name).First(&division).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "division not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve division",
		})
	}

	if err := c.BodyParser(&division); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	if err := database.DB.Save(&division).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update division",
		})
	}

	return c.JSON(fiber.NewError(fiber.StatusOK, "division updated"))
}
func DeleteDivisionByName(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "Delete Division By Name"))
}
func DeleteDivisionById(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "Delete Division By id"))
}
