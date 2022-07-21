package customer

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"lemonde.mikedelta/server/handlers"
	custservice "lemonde.mikedelta/server/services/customer"
)

func SetCustTypeRouting(rtr fiber.Router, db *gorm.DB) fiber.Router {

	rtr.Get("/GetAllCustomerTypes", func(c *fiber.Ctx) error {
		svcResult := custservice.GetAllCustomerTypes(db, c)
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	rtr.Get("/GetCustomerTypeForId", func(c *fiber.Ctx) error {
		svcResult := custservice.GetCustomerTypeForId(db, c)
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	return nil
}
