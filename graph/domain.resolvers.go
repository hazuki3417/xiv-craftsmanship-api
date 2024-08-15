package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	"github.com/hazuki3417/xiv-craftsmanship-api/graph/model"
)

// Crafts is the resolver for the crafts field.
func (r *queryResolver) Crafts(ctx context.Context, name string) ([]*model.Craft, error) {
	crafts, err := r.domain.Domain.UseCase.GetCrafts(name)

	if err != nil {
		return nil, err
	}

	result := []*model.Craft{}
	for _, craft := range crafts {
		result = append(result, &model.Craft{
			ID:     craft.ID,
			Name:   craft.Name,
			Job:    craft.Job,
			Pieces: craft.Pieces,
			Level: &model.Level{
				Item:  craft.Level.Item,
				Craft: craft.Level.Craft,
			},
		})
	}
	return result, nil
}

// Materials is the resolver for the materials field.
func (r *queryResolver) Materials(ctx context.Context, craftID string) ([]*model.Material, error) {
	materials, err := r.domain.Domain.UseCase.GetMaterials(craftID)

	if err != nil {
		return nil, err
	}

	result := []*model.Material{}
	for _, material := range materials {
		result = append(result, &model.Material{
			ParentID:   material.ParentItemId,
			ChildID:    material.ChildItemId,
			ParentName: material.ParentName,
			ChildName:  material.ChildName,
			Unit:       material.Unit,
			Total:      material.Total,
		})
	}

	return result, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
