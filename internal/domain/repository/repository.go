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
		SELECT recipe_id, item_id, name, pieces, job, item_level, recipe_level
		FROM crafts
		WHERE name ILIKE $1
		ORDER BY name
		LIMIT $2
	`

	limit := 8
	err := r.postgresql.SelectContext(ctx, &crafts, query, "%"+name+"%", limit)
	if err != nil {
		return nil, err
	}

	return crafts, nil
}

func (r *Repository) GetMaterials(craftId string) ([]*schema.Material, error) {
	ctx := context.Background()
	var materials []*schema.Material

	query := `
		SELECT
			m.tree_id,
			m.parent_item_id,
			m.parent_item_name,
			r.level as parent_craft_level,
			r.job as parent_craft_job,
			m.child_item_id,
			m.child_item_name,
			m.child_item_type,
			m.unit,
			m.total
		FROM get_materials($1) m
		JOIN recipes r
		ON m.recipe_id = r.id
	`

	err := r.postgresql.SelectContext(ctx, &materials, query, craftId)
	if err != nil {
		return nil, err
	}

	return materials, nil
}
