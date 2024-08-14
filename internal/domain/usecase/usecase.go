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
			ID:   craft.Id,
			Name: craft.Name,
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
			ParentItemId: material.ParentItemId,
			ParentName:   material.ParentName,
			ChildItemId:  material.ChildItemId,
			ChildName:    material.ChildName,
			Unit:         material.Unit,
			Total:        material.Total,
		})
	}

	return result, nil
}
