package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/xdorro/golang-fiber-project/internal/config"
)

func Middleware(a *fiber.App) {
	server := config.Server()

	a.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	// Add Recover
	a.Use(recover.New())

	// Add Icon
	a.Use(favicon.New())

	// Add CORS to each route.
	a.Use(cors.New())

	// Add simple logger.
	if server.Logger {
		a.Use(logger.New())
	}

	// Add caching.
	//cache.New(),
}
