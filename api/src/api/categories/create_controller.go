package categories

import (
	"contablue/src/domain/categories"
	"contablue/src/infra/repositories"
	"context"
	"log"
	"net/http"

	"github.com/joaocprofile/goh/core"
	"github.com/joaocprofile/goh/database/cachedb"
	"github.com/joaocprofile/goh/httpwr"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category categories.Category
	if err := httpwr.ReadBody(w, r, &category); err != nil {
		return
	}

	if err := category.Prepare(); err != nil {
		httpwr.Error(w, err.StatusCode, err.Err)
		return
	}

	categoryRepo := repositories.NewCategoryRepository()
	createCategoryService := categories.NewCreateCategoryService(categoryRepo)
	createdCategory, err := createCategoryService.Execute(&category)
	if err != nil {
		httpwr.Error(w, err.StatusCode, err.Err)
		return
	}

	redis := cachedb.NewConnection()
	if err := redis.Set(context.Background(), createdCategory.ID, createdCategory); err != nil {
		log.Println(core.Red("Error writing request to cache"), err)
	}

	httpwr.Response(w, http.StatusCreated, createdCategory)
}
