package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"lemonde.mikedelta/server/handlers"
	"lemonde.mikedelta/server/services/geoservice"
)

//Pointer to the GORM connection pool
var dbConnPool *gorm.DB

func IsAlive() {
	fmt.Println("Hello World")
}

//Initializes the REST API server and endpoint routing
func InitAPIServer() {

	//Put these in Key Vault for next demo
	dsn := "host=localhost user=lemonde_dev password=LeMonde dbname=LMIMMasterData port=5432 sslmode=disable"

	//Initialize database context for "geo" schema
	print("Initializing database context for schema geo ....")

	geoDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "geo.",
			SingularTable: true,
		},
	})

	handlers.GenericErrorHandler(err, "Cannot Connect to DB")

	print("Database Connection Pool Initialized")

	app := fiber.New()
	api := app.Group("/MikeDelta/api")

	//Create routing group and routes for Geographic information
	geo_group := api.Group("/Geographic")

	geo_group.Get("/GetAllCountries", func(c *fiber.Ctx) error {
		svcResult := geoservice.GetAllCountries(geoDB)
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	geo_group.Get("/GetCountryById", func(c *fiber.Ctx) error {
		svcResult := geoservice.GetCountryById(geoDB, c)
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	geo_group.Post("/CreateCountry", func(c *fiber.Ctx) error {
		svcResult := geoservice.CreateCountry(geoDB, string(c.Request().Body()))
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	geo_group.Delete("/DeleteCountry", func(c *fiber.Ctx) error {
		svcResult := geoservice.DeleteCountry(geoDB, c)
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	geo_group.Put("/UpdateCountry", func(c *fiber.Ctx) error {
		svcResult := geoservice.UpdateCountry(geoDB, string(c.Request().Body()))
		return handlers.GenericSvcResultHandler(svcResult, c)
	})

	app.Listen(":3000")
}
