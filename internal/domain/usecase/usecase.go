package usecase

import (
	"github.com/hazuki3417/xiv-craftsmanship-api/internal/domain/payload"
	"github.com/hazuki3417/xiv-craftsmanship-api/internal/domain/repository"
	"github.com/hazuki3417/xiv-craftsmanship-api/internal/domain/service"
	"go.uber.org/zap"
)

type UseCase struct {
	logger     *zap.Logger
	service    *service.Service
	repository *repository.Repository
}

func New(logger *zap.Logger, service *service.Service, repository *repository.Repository) *UseCase {
	return &UseCase{
		logger:     logger,
		service:    service,
		repository: repository,
	}
}

func (u *UseCase) GetCrafts(name string) ([]*payload.Craft, error) {
	crafts, err := u.repository.GetCrafts(name)
	if err != nil {
		return nil, err
	}

	var result []*payload.Craft
	for _, craft := range crafts {
		result = append(result, &payload.Craft{
			ID:   craft.ItemId,
			Name: craft.ItemName,
		})
	}

	return result, nil
}

func (u *UseCase) GetRecipe(craftId string) (*payload.Recipe, error) {
	materials, err := u.repository.GetMaterialTree(craftId)
	if err != nil {
		return nil, err
	}

	tmpNodes := make(map[string]*payload.Node)
	edges := []payload.Edge{}
	for _, material := range materials {
		if _, exists := tmpNodes[material.ChildItemId]; !exists {
			tmpNodes[material.ChildItemId] = &payload.Node{
				ID:       material.ChildItemId,
				Name:     material.ChildName,
				Unit:     material.Unit,
				Total:    material.Total,
				X:        material.X,
				Y:        material.Y,
				NodeType: material.NodeType,
			}
		}

		edges = append(edges, payload.Edge{
			Source: material.ParentItemId,
			Target: material.ChildItemId,
		})
	}

	nodes := []payload.Node{}
	for _, tmpNode := range tmpNodes {
		nodes = append(nodes, *tmpNode)
	}

	return &payload.Recipe{
		Nodes: nodes,
		Edges: edges,
	}, nil
}
