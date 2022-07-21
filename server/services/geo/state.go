package geoservice

import (
	"fmt"

	"lemonde.mikedelta/server/handlers"
	"lemonde.mikedelta/server/models/general"
	"lemonde.mikedelta/server/models/geo"

	"encoding/json"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

func GetAllStatesForCountryId(db *gorm.DB, c *fiber.Ctx) string {

	countryId, parmStatus := handlers.ParameterCheck(c, "countryid")
	if len(parmStatus) > 0 {
		return parmStatus
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

	countryId, parmStatus := handlers.ParameterCheck(c, "countryid")
	if len(parmStatus) > 0 {
		return parmStatus
	}

	stateId, parmStatus := handlers.ParameterCheck(c, "stateid")
	if len(parmStatus) > 0 {
		return parmStatus
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

	countryId, parmStatus := handlers.ParameterCheck(c, "countryid")
	if len(parmStatus) > 0 {
		return parmStatus
	}

	stateId, parmStatus := handlers.ParameterCheck(c, "stateid")
	if len(parmStatus) > 0 {
		return parmStatus
	}

	reqModel := geo.State{}

	dbResult := db.Where(map[string]interface{}{"countryid": countryId,
		"stateid": stateId}).Delete(&reqModel)
	return handlers.GenericSvcErrHandler(dbResult.Error)
}
