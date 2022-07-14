package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

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
			SingularTable: false,
		},
	})

	GenericErrorHandler(err, "Cannot Connect to DB")

	print("Database Connection Pool Initialized")

	app := fiber.New()
	api := app.Group("/MikeDelta/api")

	//Create routing group and routes for Geographic information
	geo_group := api.Group("/Geographic")

	geo_group.Get("/GetAllCountries", func(c *fiber.Ctx) error {
		return c.SendString(geoservice.GetAllCountries(geoDB))
	})

	geo_group.Get("/GetCountry", func(c *fiber.Ctx) error {
		return c.SendString("Not Implemented")
	})

	app.Listen(":3000")
}

func GenericErrorHandler(err error, msg string) {
	if err != nil {
		errmsg := fmt.Sprintf("Error! %s: %v", msg, err.Error())
		panic(errmsg)
	}
}
