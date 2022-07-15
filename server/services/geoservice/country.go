package geoservice

import (
	"fmt"

	"lemonde.mikedelta/server/models/geo"

	"gorm.io/gorm"
)

func GetAllCountries(db *gorm.DB) string {
	resultset := db.Find(&geo.Country{})
	fmt.Println(resultset.Rows())
	fmt.Println(resultset.Error)
	fmt.Println(resultset)
	return "Geo-GetCountry"
}

func CreateCountry(db *gorm.DB) string {
	db.Create(&geo.Country{
		Countryid:      505,
		Countryname:    "Stonistan",
		Iso3:           "STN",
		Iso2:           "SN",
		Countrycode:    "STN",
		Phonecode:      "99",
		Capitalcity:    "Les Paulisa",
		Currencycode:   "H",
		Internetcode:   ".sjm",
		Georegionid:    0,
		Geosubregionid: 0,
	})

	return "Created Country"
}
