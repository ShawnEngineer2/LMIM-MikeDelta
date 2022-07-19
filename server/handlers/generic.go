package handlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GenericSvcErrHandler(newErr error) string {
	if newErr != nil {
		return newErr.Error()
	} else {
		return ""
	}
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
