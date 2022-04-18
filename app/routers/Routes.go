package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/tyetis/goapiexample/app/controllers"
	"github.com/tyetis/goapiexample/app/data"
	"github.com/tyetis/goapiexample/app/services"
	"github.com/tyetis/goapiexample/app/utils"
)

var config = utils.LoadConfig()
var db, _ = data.NewConnectionDb(config)

func PostRoutes(app *fiber.App) {
	app.Get("/posts", controller.PostController.Index)
	app.Post("/post", controller.PostController.Create)
}

func UserRoutes(app *fiber.App) {
	var userService = services.NewUserService(db)
	var userController = controller.NewUserController(userService)

	app.Get("/users/:id?", userController.Index)
	app.Post("/users", userController.Create)
}

func OrderRoutes(app *fiber.App) {
	var orderService = services.NewOrderService(db)
	var orderController = controller.NewOrderController(orderService)

	app.Get("/orders/:id?", orderController.Index)
	app.Post("/orders", orderController.Create)
}

func CustomRoutes(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		err := c.SendStatus(fiber.StatusNotFound)
		return err
	})
}
