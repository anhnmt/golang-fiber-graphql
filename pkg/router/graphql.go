package router

import (
	"github.com/gofiber/fiber/v2"
)

// graphqlRoute: handler graphql query
func graphqlRoute(app *fiber.App) {
	//app.All("/query", graphqlHandler())
	//
	//app.All("/", playgroundHandler())
}

//func graphqlHandler() fiber.Handler {
//	server := config.Server()
//	client := database.Connection()
//
//	srv := handler.NewDefaultServer(resolver.NewSchema(client))
//	srv.Use(entgql.Transactioner{TxOpener: client})
//
//	if server.Logger {
//		srv.Use(&debug.Tracer{})
//	}
//
//	h := handler.NewDefaultServer(
//		generated.NewExecutableSchema(
//			generated.Config{Resolvers: &resolver.Resolver{}},
//		),
//	)
//	return adaptor.HTTPHandler(h)
//}
//
//func playgroundHandler() fiber.Handler {
//	h := playground.Handler("GraphQL", "/query")
//	return adaptor.HTTPHandler(h)
//}
