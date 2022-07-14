package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"lemonde.mikedelta/server/services/geoservice"
)

//Pointer to the GORM connection pool
var dbConnPool *gorm.DB

func IsAlive() {
	fmt.Println("Hello World")
}

//Initializes the REST API server and endpoint routing
func InitAPIServer() {

	//Initialize database connection pool
	print("Initializing database connection pool ....")

	//Put these in Key Vault for next demo
	dsn := "host=localhost user=lemonde_dev password=LeMonde dbname=LMIMMasterData port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	GenericErrorHandler(err, "Cannot Connect to DB")

	intDB, err := db.DB() //Get an interface to DB for setting the connection pool

	GenericErrorHandler(err, "Cannot Connect to DB")

	intDB.SetMaxIdleConns(10)
	intDB.SetMaxOpenConns(50)
	dbConnPool = db

	print("Database Connection Pool Initialized")

	app := fiber.New()
	api := app.Group("/MikeDelta/api")

	//Create routing group and routes for Geographic information
	geo_group := api.Group("/Geographic")

	geo_group.Get("/GetCountryList", func(c *fiber.Ctx) error {
		return c.SendString(geoservice.GetCountry())
	})

	app.Listen(":3000")
}

func GenericErrorHandler(err error, msg string) {
	if err != nil {
		errmsg := fmt.Sprintf("Error! %s: %v", msg, err.Error())
		panic(errmsg)
	}
}
