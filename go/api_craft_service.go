package openapi

import (
	"context"
	"net/http"

	"github.com/hazuki3417/xiv-craftsmanship-api/internal"
)

type CraftAPIService struct {
	service internal.Domain
}

func NewCraftAPIService(service internal.Domain) *CraftAPIService {
	return &CraftAPIService{service}
}

func (s *CraftAPIService) GetCraft(ctx context.Context, name string, body map[string]interface{}) (ImplResponse, error) {
	crafts, err := s.service.Domain.UseCase.GetCrafts(name)
	if err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}

	result := []Craft{}
	for _, craft := range crafts {
		result = append(result, Craft{
			RecipeId:   craft.Id,
			ItemId:     craft.ItemId,
			Name:       craft.Name,
			Pieces:     1,
			Job:        craft.Job,
			ItemLevel:  int32(*craft.Level.Item),
			CraftLevel: nil,
		})
	}

	return Response(http.StatusOK, result), nil
}
