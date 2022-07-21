package customer

import (
	"fmt"

	"lemonde.mikedelta/server/handlers"
	"lemonde.mikedelta/server/models/customer"

	"encoding/json"

	"gorm.io/gorm"

	"strconv"

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

	parmName := "custtypeid"

	custTypeId, _ := strconv.Atoi(c.Query(parmName, "-1"))

	//Validate input, Get a model instance, and delete
	//---------------------------------------------------------------
	if custTypeId == -1 {
		return fmt.Sprintf("Parameter Not Found: %s", parmName)
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
