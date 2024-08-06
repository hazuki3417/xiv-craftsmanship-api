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
			ID:   craft.ItemId,
			Name: craft.ItemName,
		})
	}

	return result, nil
}
