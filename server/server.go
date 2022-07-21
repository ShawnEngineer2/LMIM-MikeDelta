package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"lemonde.mikedelta/server/handlers"
	custroute "lemonde.mikedelta/server/routing/customer"
	georoute "lemonde.mikedelta/server/routing/geo"
)

func IsAlive() {
	fmt.Println("Hello World")
}

//Initializes the REST API server and endpoint routing
func InitAPIServer() {

	//Put these in Key Vault for next demo
	dsn := "host=localhost user=lemonde_dev password=LeMonde dbname=LMIMMasterData port=5432 sslmode=disable"

	//Initialize database context for "geo" schema
	print("Initializing database context for schema geo ....")

	geoDB, geoErr := handlers.ConfigDBConnection(dsn, "geo")
	handlers.GenericErrorHandler(geoErr, "Cannot Connect to schema")

	custDB, custErr := handlers.ConfigDBConnection(dsn, "customer")
	handlers.GenericErrorHandler(custErr, "Cannot Connect to schema")

	print("Database Connection Pool Initialized")

	app := fiber.New()
	api := app.Group("/MikeDelta/api")
	geo_group := api.Group("/Geographic")

	georoute.SetCountryRouting(geo_group, geoDB)
	georoute.SetStateRouting(geo_group, geoDB)
	georoute.SetCityRouting(geo_group, geoDB)

	cust_group := api.Group("/Customer")
	custroute.SetCustTypeRouting(cust_group, custDB)

	app.Listen(":3000")
}
