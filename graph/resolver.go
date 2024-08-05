package graph

import (
	"fmt"

	validator "github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	logger    *zap.Logger
	validator *validator.Validate
}

func NewResolver(logger *zap.Logger, validator *validator.Validate) *Resolver {
	if validator == nil {
		panic(fmt.Errorf("validator is nil"))
	}
	return &Resolver{
		logger:    logger,
		validator: validator,
	}
}
