package errorhandler

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func New() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {

		log.Printf("Error: %v\nStack Trace:\n%s", err, debug.Stack())
		if httpErr, ok := err.(core.HttpError); ok {
			return c.Status(httpErr.StatusCode).JSON(response.NewServerResponse(httpErr.Message, nil, httpErr.Info))
		}
		return c.Status(http.StatusInternalServerError).JSON(
			response.NewServerResponse("Oops! Something went wrong with this request", nil),
		)
	}
}
