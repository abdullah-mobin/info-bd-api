package routes

import (
	"info-bd-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	SetDivisionRoutes(v1)
	SetDistrictRoutes(v1)
}

func SetDivisionRoutes(router fiber.Router) {
	div := router.Group("/div")

	div.Get("/get-all-division", handlers.GetAllDivision).Name("get-all-division")
	div.Get("/get-division/:id", handlers.GetDivisionById).Name("get-division-by-id")
	div.Get("/get-division/:name", handlers.GetDivisionByName).Name("get-division-by-name")

	div.Post("/add-division", handlers.AddDivision).Name("add-division")

	div.Put("/update-division/:id", handlers.UpdateDivisionById).Name("update-division-by-id")
	div.Put("/update-division/:name", handlers.UpdateDivisionByName).Name("update-division-by-name")

	div.Delete("/delete-division/:id", handlers.DeleteDivisionById).Name("delete-division-by-id")
	div.Delete("/delete-division/:name", handlers.DeleteDivisionByName).Name("delete-division-by-name")
}

func SetDistrictRoutes(router fiber.Router) {
	dis := router.Group("/dis")

	dis.Get("/get-all-district", handlers.GetAllDistrict).Name("get-all-district")
	dis.Get("/get-all-district/:division", handlers.GetAllDistrictByDivision).Name("get-all-district-by-division")
	dis.Get("/get-all-district/:id", handlers.GetAllDistrictById).Name("get-all-district-by-id")
	dis.Get("/get-all-district/:name", handlers.GetAllDistrictByName).Name("get-all-district-by-name")

	dis.Post("/add-district", handlers.AddDistrict).Name("add-district")

	dis.Put("/update-district/:id", handlers.UpdateDistrictById).Name("update-district-by-id")
	dis.Put("/update-district/:name", handlers.UpdateDistrictByName).Name("update-district-by-name")

	dis.Delete("/delete-district/:id", handlers.DeleteDistrictById).Name("delete-district-by-id")
	dis.Delete("/delete-district/:name", handlers.DeleteDistrictByName).Name("delete-district-by-name")
	dis.Delete("/delete-all-district/:division", handlers.DeleteAllDistrictByDivision).Name("delete-district-by-division")
}
