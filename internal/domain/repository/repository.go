package repository

import (
	"context"
	"strconv"
	"strings"

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

func (r *Repository) GetParentItems(recipeId string) ([]*schema.ParentItem, error) {
	ctx := context.Background()
	var materials []*schema.ParentItem

	query := `
		SELECT
			distinct parent_item_id
		FROM get_materials($1)
	`

	err := r.postgresql.SelectContext(ctx, &materials, query, recipeId)
	if err != nil {
		return nil, err
	}

	return materials, nil
}

func (r *Repository) GetMaterials(itemIds []string) ([]*schema.Material, error) {
	ctx := context.Background()
	var materials []*schema.Material

	// Prepare placeholders for the IN clause
	placeholders := make([]string, len(itemIds))
	args := make([]interface{}, len(itemIds))
	for i, id := range itemIds {
		placeholders[i] = "$" + strconv.Itoa(i+1)
		args[i] = id
	}

	query := `
		SELECT
			it.id,
			it.recipe_id,
			it.parent_item_id,
			it.child_item_id,
			it.quantity,
			it.type
		FROM item_tree it
		WHERE it.parent_item_id IN (` + strings.Join(placeholders, ", ") + `)
	`

	err := r.postgresql.SelectContext(ctx, &materials, query, args...)
	if err != nil {
		return nil, err
	}

	return materials, nil
}
