package geo

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"lemonde.mikedelta/server/handlers"
	"lemonde.mikedelta/server/services/geoservice"
)

func SetStateRouting(rtr fiber.Router, db *gorm.DB) fiber.Router {

	rtr.Get("/GetAllStatesForCountryId", func(c *fiber.Ctx) error {
		svcResult := geoservice.GetAllStatesForCountryId(db, c)
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	rtr.Get("/GetStateForId", func(c *fiber.Ctx) error {
		svcResult := geoservice.GetCountryForId(db, c)
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	rtr.Post("/CreateState", func(c *fiber.Ctx) error {
		svcResult := geoservice.CreateState(db, string(c.Request().Body()))
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	rtr.Delete("/DeleteState", func(c *fiber.Ctx) error {
		svcResult := geoservice.DeleteCountry(db, c)
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	rtr.Put("/UpdateState", func(c *fiber.Ctx) error {
		svcResult := geoservice.UpdateCountry(db, string(c.Request().Body()))
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	return nil
}
