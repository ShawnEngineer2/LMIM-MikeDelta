package server

import (
	"fmt"
	"strings"

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
			SingularTable: true,
		},
	})

	GenericErrorHandler(err, "Cannot Connect to DB")

	print("Database Connection Pool Initialized")

	app := fiber.New()
	api := app.Group("/MikeDelta/api")

	//Create routing group and routes for Geographic information
	geo_group := api.Group("/Geographic")

	geo_group.Get("/GetAllCountries", func(c *fiber.Ctx) error {
		svcResult := geoservice.GetAllCountries(geoDB)
		return GenericSvcResultHandler(svcResult, c)
	})

	geo_group.Get("/GetCountryById", func(c *fiber.Ctx) error {
		svcResult := geoservice.GetCountryById(geoDB, c)
		return GenericSvcResultHandler(svcResult, c)
	})

	geo_group.Post("/CreateCountry", func(c *fiber.Ctx) error {
		svcResult := geoservice.CreateCountry(geoDB, string(c.Request().Body()))
		return GenericSvcResultHandler(svcResult, c)
	})

	geo_group.Delete("/DeleteCountry", func(c *fiber.Ctx) error {
		svcResult := geoservice.DeleteCountry(geoDB, c)
		return GenericSvcResultHandler(svcResult, c)
	})

	geo_group.Put("/UpdateCountry", func(c *fiber.Ctx) error {
		svcResult := geoservice.UpdateCountry(geoDB, string(c.Request().Body()))
		return GenericSvcResultHandler(svcResult, c)
	})

	app.Listen(":3000")
}

func GenericErrorHandler(err error, msg string) {
	if err != nil {
		errmsg := fmt.Sprintf("Error! %s: %v", msg, err.Error())
		panic(errmsg)
	}
}

func GenericSvcResultHandler(svcResult string, c *fiber.Ctx) error {

	var statusCode int = 0

	//Assign happy-path status code
	switch c.Method() {
	case "POST":
		statusCode = fiber.StatusCreated
	case "DELETE":
		statusCode = fiber.StatusNoContent
	case "GET":
		statusCode = fiber.StatusOK
	default:
		statusCode = fiber.StatusNoContent
	}

	//Handle error content
	if len(svcResult) > 0 {
		var svcResultType string = "svcUnknown"

		//Error occurred - identify and handle
		if strings.HasPrefix(svcResult, "Parameter Not Found:") {
			svcResultType = "reqError"
			statusCode = fiber.StatusBadRequest
		} else if strings.HasPrefix(svcResult, "ERROR:") {
			svcResultType = "svcError"
			statusCode = fiber.StatusInternalServerError
		} else {
			svcResultType = "svcData"
			statusCode = fiber.StatusOK
		}

		return c.Status(statusCode).JSON(fiber.Map{
			svcResultType: svcResult,
		})
	}

	return c.SendStatus(statusCode)

}
