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

func GetAllCitiesForCountryStateId(db *gorm.DB, c *fiber.Ctx) string {

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

	parmNameCountry := "countryid"
	parmNameState := "stateid"
	parmNameCity := "cityid"

	countryId, _ := strconv.Atoi(c.Query(parmNameCountry, "-1"))
	stateId, _ := strconv.Atoi(c.Query(parmNameState, "-1"))
	cityId, _ := strconv.Atoi(c.Query(parmNameCity, "-1"))

	//Validate input, Get a model instance, and delete
	//---------------------------------------------------------------
	if countryId == -1 {
		return fmt.Sprintf("Parameter Not Found: %s", parmNameCountry)
	}

	if stateId == -1 {
		return fmt.Sprintf("Parameter Not Found: %s", parmNameState)
	}

	if cityId == -1 {
		return fmt.Sprintf("Parameter Not Found: %s", parmNameCity)
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

	parmNameCountry := "countryid"
	parmNameState := "stateid"
	parmNameCity := "cityid"

	countryId, _ := strconv.Atoi(c.Query(parmNameCountry, "-1"))
	stateId, _ := strconv.Atoi(c.Query(parmNameState, "-1"))
	cityId, _ := strconv.Atoi(c.Query(parmNameCity, "-1"))

	//Validate input, Get a model instance, and delete
	//---------------------------------------------------------------
	if countryId == -1 {
		return fmt.Sprintf("Parameter Not Found: %s", parmNameCountry)
	}

	if stateId == -1 {
		return fmt.Sprintf("Parameter Not Found: %s", parmNameState)
	}

	if cityId == -1 {
		return fmt.Sprintf("Parameter Not Found: %s", parmNameCity)
	}

	reqModel := geo.City{}

	dbResult := db.Where(map[string]interface{}{"countryid": countryId,
		"stateid": stateId, "cityid": cityId}).Delete(&reqModel)
	return handlers.GenericSvcErrHandler(dbResult.Error)
}
