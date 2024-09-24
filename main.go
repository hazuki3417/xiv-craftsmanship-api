package main

import (
	"log"
	"net/http"

	openapi "github.com/hazuki3417/xiv-craftsmanship-api/go"
	"github.com/hazuki3417/xiv-craftsmanship-api/handlefunc"
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
	router := openapi.NewRouter(
		openapi.NewCraftAPIController(openapi.NewCraftAPIService(domain)),
		openapi.NewRecipeAPIController(openapi.NewRecipeAPIService(domain)),
	)

	// NOTE: サービスとは別に必要なエンドポイントは個別に追加
	router.HandleFunc("/health", handlefunc.GetHealth)
	router.HandleFunc("/openapi.yaml", handlefunc.GetOpenApi)

	log.Fatal(http.ListenAndServe(":"+container.Env.Port, router))
}
