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

func GetAllCountries(db *gorm.DB) string {
	reqModel := []geo.Country{}

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

func GetCountryForId(db *gorm.DB, c *fiber.Ctx) string {

	countryId, parmStatus := handlers.ParameterCheck(c, "countryid")
	if len(parmStatus) > 0 {
		return parmStatus
	}

	reqModel := geo.Country{}

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

func UpdateCountry(db *gorm.DB, inputData string) string {

	reqModel := geo.Country{}
	reqJson := general.DataGram01{}
	json.Unmarshal([]byte(inputData), &reqJson)

	dbResult := db.Model(&reqModel).Where(fmt.Sprintf("%s = ?", reqJson.RecKeyColumn), reqJson.RecKeyValue).Update(reqJson.UpdateColumn, reqJson.UpdateColumnValue)
	return handlers.GenericSvcErrHandler(dbResult.Error)

}

func CreateCountry(db *gorm.DB, inputData string) string {

	//Try to read the incoming JSON into a Country struct instance
	//---------------------------------------------------------------
	reqJson := geo.Country{}
	json.Unmarshal([]byte(inputData), &reqJson)

	dbResult := db.Create(&reqJson)
	return handlers.GenericSvcErrHandler(dbResult.Error)
}

func DeleteCountry(db *gorm.DB, c *fiber.Ctx) string {

	countryId, parmStatus := handlers.ParameterCheck(c, "countryid")
	if len(parmStatus) > 0 {
		return parmStatus
	}

	reqModel := geo.Country{}

	dbResult := db.Where("countryid = ?", countryId).Delete(&reqModel)
	return handlers.GenericSvcErrHandler(dbResult.Error)
}
