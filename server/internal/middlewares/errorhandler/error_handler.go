package errorhandler

import (
	"fmt"
	"net/http"

	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func New() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {

		fmt.Print(err.Error())
		if httpErr, ok := err.(core.HttpError); ok {
			return c.Status(httpErr.StatusCode).JSON(response.NewServerResponse(httpErr.Message, nil, httpErr.Info))
		}
		return c.Status(http.StatusInternalServerError).JSON(
			response.NewServerResponse("Oops! Something went wrong with this request", nil),
		)
	}
}
