package geo

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"lemonde.mikedelta/server/handlers"
	"lemonde.mikedelta/server/services/geoservice"
)

func SetStateRouting(rtr fiber.Router, db *gorm.DB, apiGroupName string) fiber.Router {

	geo_group := rtr.Group(apiGroupName)

	geo_group.Get("/GetAllStatesForCountryId", func(c *fiber.Ctx) error {
		svcResult := geoservice.GetAllCountries(db)
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	geo_group.Get("/GetStateForId", func(c *fiber.Ctx) error {
		svcResult := geoservice.GetCountryForId(db, c)
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	geo_group.Post("/CreateState", func(c *fiber.Ctx) error {
		svcResult := geoservice.CreateCountry(db, string(c.Request().Body()))
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	geo_group.Delete("/DeleteState", func(c *fiber.Ctx) error {
		svcResult := geoservice.DeleteCountry(db, c)
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	geo_group.Put("/UpdateState", func(c *fiber.Ctx) error {
		svcResult := geoservice.UpdateCountry(db, string(c.Request().Body()))
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	return nil
}
