package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"lemonde.mikedelta/server/handlers"
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

	geoDB, err := handlers.ConfigDBConnection(dsn, "geo")
	handlers.GenericErrorHandler(err, "Cannot Connect to DB")

	print("Database Connection Pool Initialized")

	app := fiber.New()
	api := app.Group("/MikeDelta/api")
	geo_group := api.Group("/Geographic")

	georoute.SetCountryRouting(geo_group, geoDB)
	georoute.SetStateRouting(geo_group, geoDB)
	georoute.SetCityRouting(geo_group, geoDB)

	app.Listen(":3000")
}
