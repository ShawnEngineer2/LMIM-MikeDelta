package geo

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"lemonde.mikedelta/server/handlers"
	"lemonde.mikedelta/server/services/geoservice"
)

func SetCountryRouting(rtr fiber.Router, db *gorm.DB) fiber.Router {

	rtr.Get("/GetAllCountries", func(c *fiber.Ctx) error {
		svcResult := geoservice.GetAllCountries(db)
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	rtr.Get("/GetCountryForId", func(c *fiber.Ctx) error {
		svcResult := geoservice.GetCountryForId(db, c)
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	rtr.Post("/CreateCountry", func(c *fiber.Ctx) error {
		svcResult := geoservice.CreateCountry(db, string(c.Request().Body()))
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	rtr.Delete("/DeleteCountry", func(c *fiber.Ctx) error {
		svcResult := geoservice.DeleteCountry(db, c)
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	rtr.Put("/UpdateCountry", func(c *fiber.Ctx) error {
		svcResult := geoservice.UpdateCountry(db, string(c.Request().Body()))
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	return nil
}
