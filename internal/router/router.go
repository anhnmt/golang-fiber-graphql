package router

import (
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg": "Welcome to Fiber Go API!",
		})
	})

	// GraphQL router
	graphqlRoute(app)

	// 404 error
	notFoundRoute(app)
}
