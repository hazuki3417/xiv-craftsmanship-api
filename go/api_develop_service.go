package openapi

import (
	"context"
	"log"
	"os"
)

type DevelopAPIService struct {
}

func NewDevelopAPIService() *DevelopAPIService {
	return &DevelopAPIService{}
}

func (s *DevelopAPIService) GetHealth(ctx context.Context) (ImplResponse, error) {
	return Response(200, nil), nil
}

// GetOpenapi - openapi
func (s *DevelopAPIService) GetOpenapi(ctx context.Context) (ImplResponse, error) {
	data, err := os.ReadFile("api/openapi.yaml")
	if err != nil {
		log.Fatal(err)
	}
	return Response(200, string(data)), nil
}
