package router

import (
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/graph/resolver"
	"github.com/xdorro/golang-fiber-base-project/internal/config"
	"github.com/xdorro/golang-fiber-base-project/internal/database"
)

// graphqlRoute: handler graphql query
func graphqlRoute(app *fiber.App) {
	app.All("/query", graphqlHandler())

	app.All("/", playgroundHandler())
}

func graphqlHandler() fiber.Handler {
	server := config.Server()
	client := database.Connection()

	srv := handler.NewDefaultServer(resolver.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})

	if server.Logger {
		srv.Use(&debug.Tracer{})
	}

	return adaptor.HTTPHandler(srv)
}

func playgroundHandler() fiber.Handler {
	h := playground.Handler("GraphQL", "/query")
	return adaptor.HTTPHandler(h)
}
