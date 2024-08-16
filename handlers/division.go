package handlers

import "github.com/gofiber/fiber/v2"

func GetAllDivision(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "Welcome all division"))
}
func GetDivisionById(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "Welcome single division id call"))
}

func GetDivisionByName(c *fiber.Ctx) error {

	return c.JSON(fiber.NewError(fiber.StatusOK, "Welcome single division name call"))
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
