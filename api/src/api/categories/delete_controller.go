package categories

import (
	"contablue/src/domain/categories"
	"contablue/src/infra/repositories"
	"net/http"

	"github.com/joaocprofile/goh/httpwr"
)

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	ID, err := httpwr.Params("id", w, r)
	if err != nil {
		return
	}

	categoryRepo := repositories.NewCategoryRepository()
	deleteCategoryService := categories.NewDeleteCategoryService(categoryRepo)
	if err = deleteCategoryService.Execute(ID); err != nil {
		httpwr.Error(w, http.StatusInternalServerError, err)
		return
	}

	httpwr.Response(w, http.StatusNoContent, nil)
}
