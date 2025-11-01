package core

import (
	"fmt"
	"net/http"

	"github.com/aritradevelops/authinfinity/server/internal/middlewares/translator"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type Controller[S Schema] interface {
	List(*fiber.Ctx) error
	Create(*fiber.Ctx) error
	View(*fiber.Ctx) error
	Update(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
}

type BaseController[S Schema] struct {
	service Service[S]
}

func NewController[S Schema](service Service[S]) Controller[S] {
	return &BaseController[S]{
		service: service,
	}
}

func (bc *BaseController[S]) List(c *fiber.Ctx) error {
	fmt.Println("here im", bc.service)
	result, err := bc.service.List(c)
	if err != nil {
		return err
	}
	return c.JSON(response.NewServerResponse(translator.Localize(c, "controller.list", map[string]string{"Entity": c.Locals("module").(string)}), result.Data, result.Info))
}

func (bc *BaseController[S]) Create(c *fiber.Ctx) error {
	id, err := bc.service.Create(c)
	if err != nil {
		fmt.Printf("Error : %+v", err)
		return err
	}

	c.Status(http.StatusCreated)
	return c.JSON(
		response.NewServerResponse(
			translator.Localize(c, "controller.create", map[string]string{"Entity": c.Locals("module").(string)}),
			fiber.Map{"id": id},
		),
	)
}
func (bc *BaseController[S]) Update(c *fiber.Ctx) error {
	acknowledged, err := bc.service.Update(c)
	if err != nil {
		fmt.Printf("Error : %+v", err)
		return err
	}
	return c.JSON(
		response.NewServerResponse(
			translator.Localize(c, "controller.update", map[string]string{"Entity": c.Locals("module").(string)}),
			fiber.Map{"acknowledged": acknowledged},
		),
	)
}
func (bc *BaseController[S]) View(c *fiber.Ctx) error {
	data, err := bc.service.View(c)
	if err != nil {
		fmt.Printf("Error : %+v", err)
		return err
	}

	return c.JSON(
		response.NewServerResponse(
			translator.Localize(c, "controller.view", map[string]string{"Entity": c.Locals("module").(string)}),
			data,
		),
	)
}

func (bc *BaseController[S]) Delete(c *fiber.Ctx) error {
	data, err := bc.service.Delete(c)
	if err != nil {
		fmt.Printf("Error: %+v", err)
		return err
	}
	return c.JSON(
		response.NewServerResponse(
			translator.Localize(c, "controller.delete", map[string]string{"Entity": c.Locals("module").(string)}),
			data,
		),
	)
}
