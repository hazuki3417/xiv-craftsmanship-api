package repository

import (
	validator "github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Repository struct {
	logger    *zap.Logger
	validator *validator.Validate
}

func New(logger *zap.Logger, validator *validator.Validate) *Repository {
	return &Repository{
		logger:    logger,
		validator: validator,
	}
}
