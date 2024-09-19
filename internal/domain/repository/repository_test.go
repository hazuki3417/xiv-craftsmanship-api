package repository_test

import (
	"fmt"
	"testing"

	"github.com/hazuki3417/xiv-craftsmanship-api/internal"
	"github.com/hazuki3417/xiv-craftsmanship-api/internal/domain/repository"
)

func TestGetParentItemId(t *testing.T) {
	container, close := internal.NewContainer()
	defer close()

	repository := repository.New(container.Logger, container.Validator, container.PostgreSQL)

	ids, err := repository.GetParentItems("7058bc49df4")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(ids)
}

func TestGetMaterials(t *testing.T) {
	container, close := internal.NewContainer()
	defer close()

	repository := repository.New(container.Logger, container.Validator, container.PostgreSQL)

	ids := []string{
		"5c7873bf132",
		"63ab375b52f",
		"37508e7e63e",
		"a422d22ecdd",
		"0975c5177af",
	}
	materials, err := repository.GetMaterials(ids)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(materials)
}
