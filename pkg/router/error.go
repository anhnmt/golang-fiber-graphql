package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-project/pkg/util"
)

// notFoundRoute: Handler page not found 404
func notFoundRoute(app *fiber.App) {
	app.Use(
		func(c *fiber.Ctx) error {
			return util.ResponseNotFound("Đường dẫn không tồn tại")
		},
	)
}
