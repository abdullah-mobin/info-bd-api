package handlers

import "github.com/gofiber/fiber/v2"

func GetAllDistrict(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "Welcome district"))
}

func GetAllDistrictByDivision(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "searching districts by division"))
}

func GetAllDistrictById(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "searching districts by id"))
}

func GetAllDistrictByName(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "searching districts by name"))
}

func AddDistrict(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "district added"))
}

func UpdateDistrictByName(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "Update District By Name"))
}
func UpdateDistrictById(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "Update District By id"))
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
