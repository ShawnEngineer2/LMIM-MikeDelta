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

	georoute.SetCountryRouting(api, geoDB, "/Geographic")
	georoute.SetStateRouting(api, geoDB, "/Geographic")

	app.Listen(":3000")
}
