package openapi

import (
	"context"
	"errors"
	"net/http"

	"github.com/hazuki3417/xiv-craftsmanship-api/internal"
)

type RecipeAPIService struct {
	service *internal.Domain
}

func NewRecipeAPIService(service *internal.Domain) *RecipeAPIService {
	return &RecipeAPIService{service}
}

func (s *RecipeAPIService) GetRecipe(ctx context.Context, recipeId string, body map[string]interface{}) (ImplResponse, error) {
	// TODO - update GetRecipe with the required logic for this service method.
	// Add api_recipe_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Material{}) or use other options such as http.Ok ...
	// return Response(200, Material{}), nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetRecipe method not implemented")
}
