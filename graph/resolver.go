package graph

import (
	"fmt"

	validator "github.com/go-playground/validator/v10"
	"github.com/hazuki3417/xiv-craftsmanship-api/internal"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	logger    *zap.Logger
	validator *validator.Validate
	domain    *internal.Domain
}

func NewResolver(logger *zap.Logger, validator *validator.Validate, domain *internal.Domain) *Resolver {
	if validator == nil {
		panic(fmt.Errorf("validator is nil"))
	}
	return &Resolver{
		logger:    logger,
		validator: validator,
		domain:    domain,
	}
}
