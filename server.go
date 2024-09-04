package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hazuki3417/xiv-craftsmanship-api/graph"
	"github.com/hazuki3417/xiv-craftsmanship-api/healthcheck"
	"github.com/hazuki3417/xiv-craftsmanship-api/internal"
)

func main() {
	container, close := graph.NewContainer()
	defer close()

	domain := internal.NewDomain(
		container.Logger,
		container.Validator,
		container.PostgreSQL,
	)

	// init graphql
	resolver := graph.NewResolver(
		container.Logger,
		container.Validator,
		domain,
	)

	config := graph.Config{
		Resolvers: resolver,
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(config))

	// NOTE: playground endpoint（development only）
	if container.Env.Stage == "development" {
		http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	}
	// NOTE: graphql endpoint
	http.Handle("/graphql", srv)
	// NOTE: healthcheck endpoint
	http.HandleFunc("/health", healthcheck.HandleFunc)

	log.Fatal(http.ListenAndServe(":"+container.Env.Port, nil))
}
