package usecase

import (
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
			Id:     craft.Id,
			ItemId: craft.ItemId,
			Name:   craft.Name,
			Job:    craft.Job,
			Pieces: craft.Pieces,
			Level: payload.Level{
				Item:  craft.ItemLevel,
				Craft: craft.RecipeLevel,
			},
		})
	}

	return result, nil
}

func (u *UseCase) GetMaterials(craftId string) ([]*payload.Material, error) {
	materials, err := u.repository.GetMaterials(craftId)
	if err != nil {
		return nil, err
	}

	var result []*payload.Material
	for _, material := range materials {
		result = append(result, &payload.Material{
			TreeId: material.TreeId,
			Parent: payload.Parent{
				ItemId:     material.ParentItemId,
				ItemName:   material.ParentItemName,
				CraftJob:   material.ParentCraftJob,
				CraftLevel: material.ParentCraftLevel,
			},
			Child: payload.Child{
				ItemId:    material.ChildItemId,
				ItemName:  material.ChildItemName,
				ItemType:  material.ChildItemType,
				ItemUnit:  material.Unit,
				ItemTotal: material.Total,
			},
		})
	}

	return result, nil
}
