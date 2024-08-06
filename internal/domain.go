package internal

import (
	validator "github.com/go-playground/validator/v10"
	"github.com/hazuki3417/xiv-craftsmanship-api/internal/domain"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Domain struct {
	Domain *domain.Factory
}

func NewDomain(logger *zap.Logger, validator *validator.Validate, postgresql *sqlx.DB) *Domain {
	return &Domain{
		Domain: domain.NewFactory(logger, validator, postgresql),
	}
}
