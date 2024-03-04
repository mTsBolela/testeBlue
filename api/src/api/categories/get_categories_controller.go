package categories

import (
	"contablue/src/infra/repositories"
	"context"
	"log"
	"net/http"

	"github.com/joaocprofile/goh/core"
	"github.com/joaocprofile/goh/database/cachedb"
	"github.com/joaocprofile/goh/httpwr"
)

func GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := httpwr.Params("id", w, r)
	if err != nil {
		return
	}

	redis := cachedb.NewConnection()
	categoryCache, err := redis.Get(context.Background(), id)
	if err == nil {
		httpwr.JSON(w, http.StatusOK, categoryCache)
		return
	}

	categoryRepo := repositories.NewCategoryRepository()
	category, err := categoryRepo.Get(id)
	if err != nil {
		httpwr.Error(w, http.StatusInternalServerError, err)
		return
	}

	err = redis.Set(context.Background(), id, category)
	if err != nil {
		log.Println(core.Red("Error writing request to cache"), err)
	}

	httpwr.Response(w, http.StatusOK, category)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	filter := httpwr.Query("filter", r)

	redis := cachedb.NewConnection()
	categoryCache, err := redis.Get(context.Background(), filter)
	if err == nil {
		httpwr.JSON(w, http.StatusOK, categoryCache)
		return
	}

	categoryRepo := repositories.NewCategoryRepository()
	categories, err := categoryRepo.GetAll(filter)
	if err != nil {
		httpwr.Error(w, http.StatusInternalServerError, err)
		return
	}
	err = redis.Set(context.Background(), filter, categories)
	if err != nil {
		log.Println(core.Red("Error writing request to cache"), err)
	}

	httpwr.Response(w, http.StatusOK, categories)
}
