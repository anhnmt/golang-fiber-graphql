package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/internal/config"
	"github.com/xdorro/golang-fiber-base-project/internal/database"
	"github.com/xdorro/golang-fiber-base-project/internal/middleware"
	"github.com/xdorro/golang-fiber-base-project/internal/router"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := config.Server()

	app := fiber.New(fiber.Config{
		AppName: server.Name,

		Prefork: server.Prefork,

		// Override default error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Default 500 statusCode
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				// Override status code if fiber.Error type
				code = e.Code
			}
			// Set Content-Type: application/json; charset=utf-8
			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

			// Return statusCode with error message
			return c.Status(code).Send([]byte(err.Error()))
		},
	})

	// Config Middleware
	middleware.Middleware(app)

	// Config Router
	router.Router(app)

	// signal channel to capture system calls
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// start shutdown goroutine
	go func() {
		// capture sigterm and other system call here
		<-sigCh
		log.Printf("Close database connection")
		_ = database.Close()

		log.Printf("Shutting down server...")
		_ = app.Shutdown()
	}()

	// start http server
	if err := app.Listen(fmt.Sprintf(":%d", server.Port)); err != nil {
		log.Printf("Oops... server is not running! error: %v", err)
	}
}
