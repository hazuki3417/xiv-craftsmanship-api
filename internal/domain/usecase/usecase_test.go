package usecase_test

import (
	"fmt"
	"testing"

	"github.com/hazuki3417/xiv-craftsmanship-api/internal"
	"github.com/hazuki3417/xiv-craftsmanship-api/internal/domain/repository"
	"github.com/hazuki3417/xiv-craftsmanship-api/internal/domain/service"
	"github.com/hazuki3417/xiv-craftsmanship-api/internal/domain/usecase"
)

func TestGetMaterials(t *testing.T) {
	container, close := internal.NewContainer()
	defer close()

	service := service.New(container.Logger)
	repository := repository.New(container.Logger, container.Validator, container.PostgreSQL)
	usecase := usecase.New(container.Logger, service, repository)

	recipes, err := usecase.GetRecipe("7058bc49df4")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(recipes)
}
