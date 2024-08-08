package repository

import (
	"context"

	validator "github.com/go-playground/validator/v10"
	"github.com/hazuki3417/xiv-craftsmanship-api/internal/domain/repository/schema"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Repository struct {
	logger     *zap.Logger
	validator  *validator.Validate
	postgresql *sqlx.DB
}

func New(logger *zap.Logger, validator *validator.Validate, postgresql *sqlx.DB) *Repository {
	return &Repository{
		logger:     logger,
		validator:  validator,
		postgresql: postgresql,
	}
}

func (r *Repository) GetCrafts(name string) ([]*schema.Craft, error) {
	ctx := context.Background()
	var crafts []*schema.Craft

	query := `
		SELECT item_id, item_name
        FROM crafts
        WHERE item_name ILIKE $1
        ORDER BY item_name
        LIMIT $2
	`

	limit := 8
	err := r.postgresql.SelectContext(ctx, &crafts, query, "%"+name+"%", limit)
	if err != nil {
		return nil, err
	}

	return crafts, nil
}

func (r *Repository) GetMaterialTree(craftId string) ([]*schema.MaterialTree, error) {
	ctx := context.Background()
	var materials []*schema.MaterialTree

	query := `
		SELECT parent_item_id, child_item_id, parent_name, child_name, unit, total, depth
        FROM get_material_tree($1)
	`

	err := r.postgresql.SelectContext(ctx, &materials, query, craftId)
	if err != nil {
		return nil, err
	}

	return materials, nil
}
