package openapi

import (
	"context"
	"net/http"

	"github.com/hazuki3417/xiv-craftsmanship-api/internal"
)

type CraftAPIService struct {
	service *internal.Domain
}

func NewCraftAPIService(service *internal.Domain) *CraftAPIService {
	return &CraftAPIService{service}
}

func (s *CraftAPIService) GetCraft(ctx context.Context, name string) (ImplResponse, error) {
	crafts, err := s.service.Domain.UseCase.GetCrafts(name)
	if err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}

	result := []Craft{}
	for _, craft := range crafts {

		data := Craft{
			RecipeId:   craft.RecipeId,
			ItemId:     craft.ItemId,
			Name:       craft.Name,
			Pieces:     int32(craft.Pieces),
			Job:        craft.Job,
			CraftLevel: int32(craft.CraftLevel),
		}

		if craft.ItemLevel != nil {
			tmp := int32(*craft.ItemLevel)
			data.ItemLevel = &tmp
		}

		result = append(result, data)
	}

	return Response(http.StatusOK, result), nil
}
