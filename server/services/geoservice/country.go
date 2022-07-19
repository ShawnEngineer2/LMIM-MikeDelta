package geoservice

import (
	"fmt"

	"lemonde.mikedelta/server/models/general"
	"lemonde.mikedelta/server/models/geo"

	"encoding/json"

	"gorm.io/gorm"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllCountries(db *gorm.DB) string {
	reqModel := []geo.Country{}

	dbResult := db.Find(&reqModel)

	dbError := GenericSvcErrHandler(dbResult.Error)

	if len(dbError) > 0 {
		return dbError
	}

	jsonResult, jsonErr := json.Marshal(reqModel)

	if jsonErr != nil {
		return GenericSvcErrHandler(jsonErr)
	}

	return string(jsonResult)
}

func GetCountryById(db *gorm.DB, c *fiber.Ctx) string {

	parmName := "countryid"
	countryId, _ := strconv.Atoi(c.Query(parmName, "-1"))

	//Validate input, Get a model instance, and delete
	//---------------------------------------------------------------
	if countryId == -1 {
		return fmt.Sprintf("Parameter Not Found: %s", parmName)
	}

	reqModel := geo.Country{}

	dbResult := db.Where(map[string]interface{}{"countryid": countryId}).Find(&reqModel)

	dbError := GenericSvcErrHandler(dbResult.Error)

	if len(dbError) > 0 {
		return dbError
	}

	jsonResult, jsonErr := json.Marshal(reqModel)

	if jsonErr != nil {
		return GenericSvcErrHandler(jsonErr)
	}

	return string(jsonResult)

}

func UpdateCountry(db *gorm.DB, inputData string) string {

	reqModel := geo.Country{}
	reqJson := general.DataGram{}
	json.Unmarshal([]byte(inputData), &reqJson)

	fmt.Println(reqJson)

	dbResult := db.Model(&reqModel).Where(fmt.Sprintf("%s = ?", reqJson.RecKeyColumn), reqJson.RecKeyValue).Update(reqJson.UpdateColumn, reqJson.UpdateColumnValue)
	return GenericSvcErrHandler(dbResult.Error)

}

func CreateCountry(db *gorm.DB, inputData string) string {

	//Try to read the incoming JSON into a Country struct instance
	//---------------------------------------------------------------
	reqJson := geo.Country{}
	json.Unmarshal([]byte(inputData), &reqJson)

	dbResult := db.Create(&reqJson)
	return GenericSvcErrHandler(dbResult.Error)
}

func DeleteCountry(db *gorm.DB, c *fiber.Ctx) string {

	parmName := "countryid"
	countryId, _ := strconv.Atoi(c.Query(parmName, "-1"))

	//Validate input, Get a model instance, and delete
	//---------------------------------------------------------------
	if countryId == -1 {
		return fmt.Sprintf("Parameter Not Found: %s", parmName)
	}

	reqModel := geo.Country{}

	dbResult := db.Where("countryid = ?", countryId).Delete(&reqModel)
	return GenericSvcErrHandler(dbResult.Error)
}
