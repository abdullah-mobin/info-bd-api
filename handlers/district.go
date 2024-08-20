package handlers

import (
	"errors"
	"info-bd-api/database"
	"info-bd-api/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllDistrict(c *fiber.Ctx) error {

	var dis []model.District

	res := database.DB.Find(&dis)

	if res.Error != nil {
		return c.JSON(fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch districts"))
	}

	return c.JSON(dis)
}

func GetAllDistrictByDivisionId(c *fiber.Ctx) error {
	divisionID := c.Params("division_id")
	var division model.Division

	if err := database.DB.Preload("Districts").First(&division, "id = ?", divisionID).Error; err != nil {
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

func GetAllDistrictByDivisionName(c *fiber.Ctx) error {

	name := c.Params("division_name")
	var div model.Division
	if err := database.DB.Preload("Districts").Where("name = ?", name).First(&div).Error; err != nil {
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

func GetDistrictById(c *fiber.Ctx) error {

	id := c.Params("id")
	var district model.District

	if err := database.DB.First(&district, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"erros": "District not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(district)
}

func GetDistrictByName(c *fiber.Ctx) error {

	name := c.Params("name")
	var district model.District
	if err := database.DB.Where("name = ?", name).First(&district).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "district not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch district",
		})
	}
	return c.Status(fiber.StatusOK).JSON(district)
}

func AddDistrict(c *fiber.Ctx) error {

	var district model.District
	if err := c.BodyParser(&district); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse request",
		})
	}
	if err := database.DB.Create(&district).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to save data",
		})
	}

	return c.JSON(fiber.NewError(fiber.StatusOK, "district added"))
}

func UpdateDistrictByName(c *fiber.Ctx) error {
	name := c.Params("name")
	var district model.District

	if err := database.DB.Where("name = ?", name).First(&district).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "district not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve district",
		})
	}

	if err := c.BodyParser(&district); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	if err := database.DB.Save(&district).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update district",
		})
	}

	return c.JSON(fiber.NewError(fiber.StatusOK, "district updated"))
}
func UpdateDistrictById(c *fiber.Ctx) error {
	id := c.Params("id")
	var district model.District

	if err := database.DB.First(&district, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "district not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve district",
		})
	}

	if err := c.BodyParser(&district); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	if err := database.DB.Save(&district).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update district",
		})
	}

	return c.JSON(fiber.NewError(fiber.StatusOK, "district updated"))
}

func DeleteDistrictById(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "delete districts by id"))
}
func DeleteDistrictByName(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "delete districts by name"))
}
func DeleteAllDistrictByDivision(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "delete all districts by division"))
}
