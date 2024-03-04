package persons

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
	personCache, err := redis.Get(context.Background(), id)
	if err == nil {
		httpwr.JSON(w, http.StatusOK, personCache)
		return
	}

	personRepo := repositories.NewPersonRepository()
	person, err := personRepo.Get(id)
	if err != nil {
		httpwr.Error(w, http.StatusInternalServerError, err)
		return
	}
	err = redis.Set(context.Background(), id, person)
	if err != nil {
		log.Println(core.Red("Error writing Search Person ID to cache"), err)
	}

	httpwr.Response(w, http.StatusOK, person)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	name := httpwr.Query("name", r)

	redis := cachedb.NewConnection()
	personCache, err := redis.Get(context.Background(), name)
	if err == nil {
		httpwr.JSON(w, http.StatusOK, personCache)
		return
	}

	personRepo := repositories.NewPersonRepository()
	personList, err := personRepo.GetAll(name)
	if err != nil {
		httpwr.Error(w, http.StatusInternalServerError, err)
		return
	}
	err = redis.Set(context.Background(), name, personList)
	if err != nil {
		log.Println(core.Red("Error writing request to cache"), err)
	}

	httpwr.Response(w, http.StatusOK, personList)
}
