package usecase

import (
	"errors"
	"fmt"

	"github.com/hazuki3417/xiv-craftsmanship-api/internal/domain/payload"
	"github.com/hazuki3417/xiv-craftsmanship-api/internal/domain/repository"
	"github.com/hazuki3417/xiv-craftsmanship-api/internal/domain/service"
	"go.uber.org/zap"
)

type UseCase struct {
	logger     *zap.Logger
	service    *service.Service
	repository *repository.Repository
}

func New(logger *zap.Logger, service *service.Service, repository *repository.Repository) *UseCase {
	return &UseCase{
		logger:     logger,
		service:    service,
		repository: repository,
	}
}

func (u *UseCase) GetCrafts(name string) ([]*payload.Craft, error) {
	crafts, err := u.repository.GetCrafts(name)
	if err != nil {
		return nil, err
	}

	var result []*payload.Craft
	for _, craft := range crafts {
		result = append(result, &payload.Craft{
			RecipeId:   craft.Id,
			ItemId:     craft.ItemId,
			Name:       craft.Name,
			Job:        craft.Job,
			Pieces:     craft.Pieces,
			ItemLevel:  *craft.ItemLevel,
			CraftLevel: craft.RecipeLevel,
		})
	}

	return result, nil
}

func (u *UseCase) GetRecipe(recipeId string) (*payload.Recipe, error) {
	items, err := u.repository.GetParentItems(recipeId)
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, errors.New("recipe data not found")
	}

	ids := make([]string, len(items))
	for i, parentItem := range items {
		ids[i] = parentItem.Id
	}

	materials, err := u.repository.GetMaterials(ids)
	if err != nil {
		return nil, err
	}

	crafts, err := u.repository.GetCraftsByItemIds(ids)
	if err != nil {
		return nil, err
	}

	craftMap := make(map[string]payload.Craft)
	for _, craft := range crafts {
		craftMap[craft.Id] = payload.Craft{
			RecipeId: craft.Id,
			ItemId:   craft.ItemId,
			Name:     craft.Name,
			Job:      craft.Job,
			Pieces:   craft.Pieces,
		}
	}

	// NOTE: アイテムごとのレシピリストを作成
	// item - recipe1
	//      - recipe2
	itemMap := make(map[string]map[string][]payload.Material)
	for _, material := range materials {
		if _, exists := itemMap[material.ParentItemId]; !exists {
			itemMap[material.ParentItemId] = make(map[string][]payload.Material)
		}

		itemMap[material.ParentItemId][material.RecipeId] = append(itemMap[material.ParentItemId][material.RecipeId], payload.Material{
			ItemId:   material.ChildItemId,
			ItemName: material.ChildItemName,
			Quantity: material.Quantity,
			Type:     material.Type,
			Recipes:  []payload.Recipe{},
		})
	}

	rootItemId := ""
	for _, material := range materials {
		if recipeId == material.RecipeId {
			rootItemId = material.ParentItemId
			break
		}
	}

	root := payload.Recipe{
		ItemId:   rootItemId,
		RecipeId: recipeId,
	}

	tree, err := createRecipe(root, itemMap, craftMap)
	if err != nil {
		return nil, err
	}

	return tree, nil
}

func createRecipe(
	source payload.Recipe,
	itemMap map[string]map[string][]payload.Material,
	craftMap map[string]payload.Craft,
) (*payload.Recipe, error) {

	materials, existsMaterials := itemMap[source.ItemId][source.RecipeId]
	craft, existsCraft := craftMap[source.RecipeId]

	if !existsMaterials {
		// NOTE: レシピが存在しない場合はスキップ
		return nil, nil
	}

	// レシピが存在するときの処理
	if !existsCraft {
		return nil, fmt.Errorf("craft data not found for itemId: %s", source.ItemId)
	}

	source.Pieces = craft.Pieces

	for i := range materials {
		material, err := createMaterial(materials[i], itemMap, craftMap)

		if err != nil {
			return nil, err
		}

		source.Materials = append(source.Materials, *material)
	}
	return &source, nil
}

func createMaterial(
	source payload.Material,
	itemMap map[string]map[string][]payload.Material,
	craftMap map[string]payload.Craft,
) (*payload.Material, error) {
	recipes, exists := itemMap[source.ItemId]

	if !exists {
		return &source, nil
	}

	if recipes == nil {
		return nil, errors.New("map data is nil")
	}

	for recipeId := range recipes {
		base := payload.Recipe{
			ItemId:   source.ItemId,
			RecipeId: recipeId,
		}
		recipe, err := createRecipe(base, itemMap, craftMap)

		if err != nil {
			return nil, err
		}

		source.Recipes = append(
			source.Recipes,
			*recipe,
		)
	}

	return &source, nil
}
