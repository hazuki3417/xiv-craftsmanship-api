package openapi

import (
	"context"
	"errors"
	"net/http"

	"github.com/hazuki3417/xiv-craftsmanship-api/internal"
	"github.com/hazuki3417/xiv-craftsmanship-api/internal/domain/payload"
)

type RecipeAPIService struct {
	service *internal.Domain
}

func NewRecipeAPIService(service *internal.Domain) *RecipeAPIService {
	return &RecipeAPIService{service}
}

func (s *RecipeAPIService) GetRecipe(ctx context.Context, recipeId string) (ImplResponse, error) {
	materials, err := s.service.Domain.UseCase.GetMaterials(recipeId)
	if materials == nil || len(materials.Recipes) == 0 {
		return Response(http.StatusNotFound, nil), errors.New("recipe not found")
	}

	if err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}

	tree := mapMaterialStruct(materials)

	return Response(200, tree), nil
}

func mapMaterialStruct(source *payload.Material) *Material {
	recipes := []Recipe{}

	for _, recipe := range source.Recipes {
		recipes = append(recipes, *mapRecipeStruct(&recipe))
	}

	return &Material{
		ItemId:   source.ItemId,
		Quantity: int32(source.Quantity),
		Type:     ItemType(source.Type),
		Recipes:  recipes,
	}
}

func mapRecipeStruct(source *payload.Recipe) *Recipe {
	// 元データの構造体を変換後の構造体にマッピング
	materials := []Material{}
	for _, material := range source.Materials {
		materials = append(materials, *mapMaterialStruct(&material))
	}

	return &Recipe{
		RecipeId:  source.RecipeId,
		ItemId:    source.ItemId,
		Materials: materials,
	}
}
