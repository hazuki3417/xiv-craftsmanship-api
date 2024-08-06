package domain

import (
	validator "github.com/go-playground/validator/v10"
	"github.com/hazuki3417/xiv-craftsmanship-api/internal/domain/repository"
	"github.com/hazuki3417/xiv-craftsmanship-api/internal/domain/service"
	"github.com/hazuki3417/xiv-craftsmanship-api/internal/domain/usecase"
	"go.uber.org/zap"
)

type Factory struct {
	UseCase *usecase.UseCase
}

func NewFactory(logger *zap.Logger, validator *validator.Validate) *Factory {
	repository := repository.New(logger, validator)
	service := service.New(logger)
	usecase := usecase.New(logger, service, repository)
	return &Factory{
		UseCase: usecase,
	}
}
