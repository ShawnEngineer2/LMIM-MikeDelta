package geoservice

import (
	"lemonde.mikedelta/server/models/geo"

	"gorm.io/gorm"
)

func GetAllCountries(*gorm.DB) string {
	var city = geo.City{}
	city.Cityid = 1
	return "Geo-GetCountry"
}
