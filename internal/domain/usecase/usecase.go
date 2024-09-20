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

func (u *UseCase) GetMaterials(recipeId string) (*payload.Material, error) {
	parentItems, err := u.repository.GetParentItems(recipeId)
	if err != nil {
		return nil, err
	}

	if len(parentItems) == 0 {
		return nil, errors.New("recipe data not found")
	}

	var parentItemIds []string
	for _, parentItem := range parentItems {
		parentItemIds = append(parentItemIds, parentItem.Id)
	}

	// item idを起点で再帰的にツリー構造を作成
	materials, err := u.repository.GetMaterials(parentItemIds)
	if err != nil {
		return nil, err
	}

	recipeMap := make(map[string]map[string][]payload.Material)

	// NOTE: nodeを作成
	rootItemId := ""
	for _, material := range materials {
		if _, exists := recipeMap[material.ParentItemId]; !exists {
			recipeMap[material.ParentItemId] = make(map[string][]payload.Material)
		}

		if rootItemId == "" && material.RecipeId == recipeId {
			rootItemId = material.ParentItemId
		}

		recipeMap[material.ParentItemId][material.RecipeId] = append(recipeMap[material.ParentItemId][material.RecipeId], payload.Material{
			ItemID:   material.ChildItemId,
			Quantity: material.Quantity,
			Type:     material.Type,
			Recipes:  []payload.Recipe{},
		})
	}

	tree := createTree(rootItemId, recipeMap)

	return &tree, nil
}

func createTree(itemID string, recipeMap map[string]map[string][]payload.Material) payload.Material {
	node := payload.Material{
		ItemID:   itemID,
		Quantity: 1,
		Type:     "material",
		Recipes:  []payload.Recipe{},
	}

	// レシピが存在する場合のみ処理
	if recipes, exists := recipeMap[itemID]; exists {
		for recipeID, materials := range recipes {
			recipe := payload.Recipe{
				RecipeID:  recipeID,
				ItemID:    itemID,
				Materials: []payload.Material{},
			}
			for _, material := range materials {
				// 再帰的に子ノードを作成
				childNode := createTree(material.ItemID, recipeMap)
				material.Recipes = childNode.Recipes
				recipe.Materials = append(recipe.Materials, material)
			}
			node.Recipes = append(node.Recipes, recipe)
		}
	}
	return node
}
