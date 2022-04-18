package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tyetis/goapiexample/app/data/entity"
	"github.com/tyetis/goapiexample/app/services"
)

type orderController struct {
	service services.IOrderService
}

func NewOrderController(service services.IOrderService) orderController {
	return orderController{
		service: service,
	}
}

func (c *orderController) Create(ctx *fiber.Ctx) error {
	return ctx.JSON("Create")
}

func (c *orderController) Index(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	orders := c.service.GetAll(&entity.Order{Id: id})

	return ctx.JSON(fiber.Map{
		"Orders": orders,
	})
}
