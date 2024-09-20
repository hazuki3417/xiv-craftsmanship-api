package main

import (
	"log"
	"net/http"

	openapi "github.com/hazuki3417/xiv-craftsmanship-api/go"
	"github.com/hazuki3417/xiv-craftsmanship-api/internal"
)

func main() {
	container, close := internal.NewContainer()
	defer close()

	domain := internal.NewDomain(
		container.Logger,
		container.Validator,
		container.PostgreSQL,
	)

	// NOTE: 手動で追加すること
	CraftAPIService := openapi.NewCraftAPIService(domain)
	CraftAPIController := openapi.NewCraftAPIController(CraftAPIService)

	router := openapi.NewRouter(CraftAPIController)

	log.Fatal(http.ListenAndServe(":"+container.Env.Port, router))
}
