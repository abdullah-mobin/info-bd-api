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
	div.Get("/get-division/id/:id", handlers.GetDivisionById).Name("get-division-by-id")
	div.Get("/get-division/name/:name", handlers.GetDivisionByName).Name("get-division-by-name")

	div.Post("/add-division", handlers.AddDivision).Name("add-division")

	div.Put("/update-division/id/:id", handlers.UpdateDivisionById).Name("update-division-by-id")
	div.Put("/update-division/name/:name", handlers.UpdateDivisionByName).Name("update-division-by-name")

	div.Delete("/delete-division/id/:id", handlers.DeleteDivisionById).Name("delete-division-by-id")
	div.Delete("/delete-division/name/:name", handlers.DeleteDivisionByName).Name("delete-division-by-name")
}

func SetDistrictRoutes(router fiber.Router) {
	dis := router.Group("/dis")

	dis.Get("/get-all-district", handlers.GetAllDistrict).Name("get-all-district")
	dis.Get("/get-all-district/division-name/:division_name", handlers.GetAllDistrictByDivisionName).Name("get-all-district-by-division-name")
	dis.Get("/get-all-district/division-id/:division_id", handlers.GetAllDistrictByDivisionId).Name("get-all-district-by-division-id")
	dis.Get("/get-district/id/:id", handlers.GetDistrictById).Name("get-district-by-id")
	dis.Get("/get-district/name/:name", handlers.GetDistrictByName).Name("get-district-by-name")

	dis.Post("/add-district", handlers.AddDistrict).Name("add-district")

	dis.Put("/update-district/id/:id", handlers.UpdateDistrictById).Name("update-district-by-id")
	dis.Put("/update-district/name/:name", handlers.UpdateDistrictByName).Name("update-district-by-name")

	dis.Delete("/delete-district/id/:id", handlers.DeleteDistrictById).Name("delete-district-by-id")
	dis.Delete("/delete-district/name/:name", handlers.DeleteDistrictByName).Name("delete-district-by-name")
	dis.Delete("/delete-all-district/division/:division", handlers.DeleteAllDistrictByDivision).Name("delete-district-by-division")
}
