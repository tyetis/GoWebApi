package app

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	routes "github.com/tyetis/goapiexample/app/routers"
)

func CreateServer(port int) {
	// Create Fiber App
	app := fiber.New()

	// app.Use(middleware.Example)
	routes.PostRoutes(app)
	routes.UserRoutes(app)
	routes.OrderRoutes(app)
	routes.CustomRoutes(app)

	// Start server
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
