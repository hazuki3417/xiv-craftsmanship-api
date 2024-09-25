package usecase

import (
	"errors"

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

	tree, err := createTree(rootItemId, itemMap, craftMap)
	if err != nil {
		return nil, err
	}

	return tree, nil
}

func createTree(
	currentItemId string,
	itemMap map[string]map[string][]payload.Material,
	craftMap map[string]payload.Craft,
) (*payload.Recipe, error) {

	// レシピが存在する場合のみ処理
	if recipeMap, exists := itemMap[currentItemId]; exists {

		if len(recipeMap) == 0 {
			return nil, nil
		}

		var current *payload.Recipe
		for recipeId, materials := range recipeMap {
			var craft *payload.Craft
			if c, exists := craftMap[recipeId]; !exists {
				return nil, errors.New("craft data not found")
			} else {
				craft = &c
			}

			current = &payload.Recipe{
				RecipeId:  craft.RecipeId,
				Pieces:    craft.Pieces,
				ItemId:    currentItemId,
				Materials: []payload.Material{},
			}

			for i, material := range materials {

				// 再帰的に子ノードを作成
				tree, err := createTree(material.ItemId, itemMap, craftMap)

				if err != nil {
					return nil, err
				}

				if tree == nil {
					continue
				}
				materials[i].Recipes = append(materials[i].Recipes, *tree)
			}
			current.Materials = append(current.Materials, materials...)
		}
		return current, nil
	}

	return nil, nil
}
