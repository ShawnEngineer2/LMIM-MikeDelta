package geoservice

import (
	"fmt"

	"lemonde.mikedelta/server/handlers"
	"lemonde.mikedelta/server/models/general"
	"lemonde.mikedelta/server/models/geo"

	"encoding/json"

	"gorm.io/gorm"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllStatesForCountryId(db *gorm.DB, c *fiber.Ctx) string {

	parmName := "countryid"
	countryId, _ := strconv.Atoi(c.Query(parmName, "-1"))

	//Validate input, Get a model instance, and delete
	//---------------------------------------------------------------
	if countryId == -1 {
		return fmt.Sprintf("Parameter Not Found: %s", parmName)
	}

	reqModel := []geo.State{}

	dbResult := db.Where(map[string]interface{}{"countryid": countryId}).Find(&reqModel)

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

func GetStateForCountryStateId(db *gorm.DB, c *fiber.Ctx) string {

	parmNameCountry := "countryid"
	parmNameState := "stateid"

	countryId, _ := strconv.Atoi(c.Query(parmNameCountry, "-1"))
	stateId, _ := strconv.Atoi(c.Query(parmNameState, "-1"))

	//Validate input, Get a model instance, and delete
	//---------------------------------------------------------------
	if countryId == -1 {
		return fmt.Sprintf("Parameter Not Found: %s", parmNameCountry)
	}

	if stateId == -1 {
		return fmt.Sprintf("Parameter Not Found: %s", parmNameState)
	}

	reqModel := geo.State{}

	dbResult := db.Where(map[string]interface{}{"countryid": countryId,
		"stateid": stateId}).Find(&reqModel)

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

func UpdateState(db *gorm.DB, inputData string) string {

	reqModel := geo.State{}
	reqJson := general.DataGram02{}
	json.Unmarshal([]byte(inputData), &reqJson)

	dbResult := db.Model(&reqModel).Where(fmt.Sprintf("%s = ? and %s = ?", reqJson.RecKeyColumn1, reqJson.RecKeyColumn2), reqJson.RecKeyValue1, reqJson.RecKeyValue2).Update(reqJson.UpdateColumn, reqJson.UpdateColumnValue)
	return handlers.GenericSvcErrHandler(dbResult.Error)

}

func CreateState(db *gorm.DB, inputData string) string {

	//Try to read the incoming JSON into a State struct instance
	//---------------------------------------------------------------
	fmt.Println("Boom")
	reqJson := geo.State{}
	json.Unmarshal([]byte(inputData), &reqJson)

	dbResult := db.Create(&reqJson)
	return handlers.GenericSvcErrHandler(dbResult.Error)
}

func DeleteState(db *gorm.DB, c *fiber.Ctx) string {

	parmNameCountry := "countryid"
	parmNameState := "stateid"

	countryId, _ := strconv.Atoi(c.Query(parmNameCountry, "-1"))
	stateId, _ := strconv.Atoi(c.Query(parmNameState, "-1"))

	//Validate input, Get a model instance, and delete
	//---------------------------------------------------------------
	if countryId == -1 {
		return fmt.Sprintf("Parameter Not Found: %s", parmNameCountry)
	}

	if stateId == -1 {
		return fmt.Sprintf("Parameter Not Found: %s", parmNameState)
	}

	reqModel := geo.State{}

	dbResult := db.Where(map[string]interface{}{"countryid": countryId,
		"stateid": stateId}).Delete(&reqModel)
	return handlers.GenericSvcErrHandler(dbResult.Error)
}
