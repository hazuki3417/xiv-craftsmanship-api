package usecase

import (
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
