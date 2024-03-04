package categories

import (
	"contablue/src/domain/categories"
	"contablue/src/infra/repositories"
	"errors"
	"net/http"

	"github.com/joaocprofile/goh/httpwr"
)

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	ID, err := httpwr.Params("id", w, r)
	if err != nil {
		return
	}

	var category categories.Category
	if err = httpwr.ReadBody(w, r, &category); err != nil {
		return
	}
	if err = category.Prepare(); err != nil {
		httpwr.Error(w, http.StatusBadRequest, errors.New("Error Validate Category: "+err.Error()))
		return
	}
	category.ID = ID

	categoryRepo := repositories.NewCategoryRepository()
	updateCategoryService := categories.NewUpdateCategoryService(categoryRepo)
	if err := updateCategoryService.Execute(&category); err != nil {
		httpwr.Error(w, err.StatusCode, err.Err)
		return
	}

	httpwr.Response(w, http.StatusNoContent, nil)
}
