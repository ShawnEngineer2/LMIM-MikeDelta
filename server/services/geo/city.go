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

func GetAllCitiesForCountryStateId(db *gorm.DB, c *fiber.Ctx) string {

	countryId, parmStatus := handlers.ParameterCheck(c, "countryid")
	if len(parmStatus) > 0 {
		return parmStatus
	}

	stateId, parmStatus := handlers.ParameterCheck(c, "stateid")
	if len(parmStatus) > 0 {
		return parmStatus
	}

	reqModel := []geo.City{}

	dbResult := db.Where(map[string]interface{}{"countryid": countryId, "stateid": stateId}).Find(&reqModel)

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

func GetCityForCountryStateCityId(db *gorm.DB, c *fiber.Ctx) string {

	countryId, parmStatus := handlers.ParameterCheck(c, "countryid")
	if len(parmStatus) > 0 {
		return parmStatus
	}

	stateId, parmStatus := handlers.ParameterCheck(c, "stateid")
	if len(parmStatus) > 0 {
		return parmStatus
	}

	cityId, parmStatus := handlers.ParameterCheck(c, "cityid")
	if len(parmStatus) > 0 {
		return parmStatus
	}

	reqModel := geo.City{}

	dbResult := db.Where(map[string]interface{}{"countryid": countryId,
		"stateid": stateId, "cityid": cityId}).Find(&reqModel)

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

func UpdateCity(db *gorm.DB, inputData string) string {

	reqModel := geo.City{}
	reqJson := general.DataGram03{}
	json.Unmarshal([]byte(inputData), &reqJson)

	dbResult := db.Model(&reqModel).Where(fmt.Sprintf("%s = ? and %s = ? and %s = ?", reqJson.RecKeyColumn1, reqJson.RecKeyColumn2, reqJson.RecKeyColumn3), reqJson.RecKeyValue1, reqJson.RecKeyValue2, reqJson.RecKeyValue3).Update(reqJson.UpdateColumn, reqJson.UpdateColumnValue)
	return handlers.GenericSvcErrHandler(dbResult.Error)
}

func CreateCity(db *gorm.DB, inputData string) string {

	//Try to read the incoming JSON into a State struct instance
	//---------------------------------------------------------------
	reqJson := geo.City{}
	json.Unmarshal([]byte(inputData), &reqJson)

	dbResult := db.Create(&reqJson)
	return handlers.GenericSvcErrHandler(dbResult.Error)
}

func DeleteCity(db *gorm.DB, c *fiber.Ctx) string {

	countryId, parmStatus := handlers.ParameterCheck(c, "countryid")
	if len(parmStatus) > 0 {
		return parmStatus
	}

	stateId, parmStatus := handlers.ParameterCheck(c, "stateid")
	if len(parmStatus) > 0 {
		return parmStatus
	}

	cityId, parmStatus := handlers.ParameterCheck(c, "cityid")
	if len(parmStatus) > 0 {
		return parmStatus
	}

	reqModel := geo.City{}

	dbResult := db.Where(map[string]interface{}{"countryid": countryId,
		"stateid": stateId, "cityid": cityId}).Delete(&reqModel)
	return handlers.GenericSvcErrHandler(dbResult.Error)
}
