package categories_test

import (
	"contablue/src/domain/categories"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/google/uuid"
)

func TestCategory(t *testing.T) {
	const errorDefult = "Expected response is [%v], but responsed value was [%v]"

	t.Run("when all params are valid, return a valid category struct", func(t *testing.T) {
		category := categories.NewCategory()
		category.ID = uuid.New().String()
		category.Description = "Pagamento de fornecedor"
		category.Type = "credit"
		category.Status = "active"
		require.Nil(t, category.Prepare())
	})

	t.Run("when param description is empty, returns a error", func(t *testing.T) {
		category := categories.Category{
			ID:          uuid.New().String(),
			Description: "",
			Type:        "credit",
			Status:      "active",
		}
		require.NotNil(t, category.Prepare())
		require.Error(t, category.Prepare())
	})

	t.Run("when type is empty, returns a error", func(t *testing.T) {
		category := categories.Category{
			ID:          uuid.New().String(),
			Description: "Pamento de Energia",
			Type:        "",
			Status:      "active",
		}
		require.NotNil(t, category.Prepare())
		require.Error(t, category.Prepare())
	})

	t.Run("when invalid type, returns a error", func(t *testing.T) {
		category := categories.Category{
			ID:          uuid.New().String(),
			Description: "Pamento de Energia",
			Type:        "blabla",
			Status:      "active",
		}
		require.NotNil(t, category.Prepare())
		require.Error(t, category.Prepare())
	})

	t.Run("when invalid status, returns a error", func(t *testing.T) {
		category := categories.Category{
			ID:          uuid.New().String(),
			Description: "Pamento de Energia",
			Type:        "debit",
			Status:      "",
		}
		require.NotNil(t, category.Prepare())
		require.Error(t, category.Prepare())
	})

}
