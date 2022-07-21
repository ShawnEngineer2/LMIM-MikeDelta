package customer

import (
	"lemonde.mikedelta/server/handlers"
	"lemonde.mikedelta/server/models/customer"

	"encoding/json"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

func GetAllCustomerTypes(db *gorm.DB, c *fiber.Ctx) string {

	reqModel := []customer.Customertype{}

	dbResult := db.Find(&reqModel)

	dbError := handlers.GenericSvcErrHandler(dbResult.Error)

	if len(dbError) > 0 {
		return dbError
	}

	jsonResult, jsonErr := json.Marshal(reqModel)

	if jsonErr != nil {
		return handlers.GenericSvcErrHandler(jsonErr)
	}

	return string(jsonResult)
}

func GetCustomerTypeForId(db *gorm.DB, c *fiber.Ctx) string {

	custTypeId, parmStatus := handlers.ParameterCheck(c, "custtypeid")
	if len(parmStatus) > 0 {
		return parmStatus
	}

	reqModel := customer.Customertype{}

	dbResult := db.Where(map[string]interface{}{"custtypeid": custTypeId}).Find(&reqModel)

	dbError := handlers.GenericSvcErrHandler(dbResult.Error)

	if len(dbError) > 0 {
		return dbError
	}

	jsonResult, jsonErr := json.Marshal(reqModel)

	if jsonErr != nil {
		return handlers.GenericSvcErrHandler(jsonErr)
	}

	return string(jsonResult)

}
