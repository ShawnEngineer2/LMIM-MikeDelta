package geo

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"lemonde.mikedelta/server/handlers"
	"lemonde.mikedelta/server/services/geoservice"
)

func SetCityRouting(rtr fiber.Router, db *gorm.DB) fiber.Router {

	rtr.Get("/GetAllCitiesForId", func(c *fiber.Ctx) error {
		svcResult := geoservice.GetAllCitiesForCountryStateId(db, c)
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	rtr.Get("/GetCityForId", func(c *fiber.Ctx) error {
		svcResult := geoservice.GetCityForCountryStateCityId(db, c)
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	rtr.Post("/CreateCity", func(c *fiber.Ctx) error {
		svcResult := geoservice.CreateCity(db, string(c.Request().Body()))
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	rtr.Delete("/DeleteCity", func(c *fiber.Ctx) error {
		svcResult := geoservice.DeleteCity(db, c)
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	rtr.Put("/UpdateCity", func(c *fiber.Ctx) error {
		svcResult := geoservice.UpdateCity(db, string(c.Request().Body()))
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	return nil
}
