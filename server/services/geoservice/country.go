package geoservice

import (
	"lemonde.mikedelta/server/models/geo"
)

func GetCountry() string {
	var city = geo.City{}
	city.Cityid = 1
	return "Geo-GetCountry"
}
