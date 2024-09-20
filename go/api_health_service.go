package openapi

import (
	"context"
)

type HealthAPIService struct {
}

func NewHealthAPIService() *HealthAPIService {
	return &HealthAPIService{}
}

// GetHealth - health
func (s *HealthAPIService) GetHealth(ctx context.Context, body map[string]interface{}) (ImplResponse, error) {
	return Response(200, nil), nil
}
