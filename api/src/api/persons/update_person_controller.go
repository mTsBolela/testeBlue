package persons

import (
	"contablue/src/domain/persons"
	"contablue/src/infra/repositories"
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/joaocprofile/goh/core"
	"github.com/joaocprofile/goh/database/cachedb"
	"github.com/joaocprofile/goh/httpwr"
)

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	ID, err := httpwr.Params("id", w, r)
	if err != nil {
		return
	}

	var person persons.Person
	if err := httpwr.ReadBody(w, r, &person); err != nil {
		return
	}
	if err := person.Prepare(); err != nil {
		httpwr.Error(w, http.StatusBadRequest, errors.New("Error Validate Person: "+err.Error()))
		return
	}
	person.ID = ID

	personRepo := repositories.NewPersonRepository()
	updatePersonService := persons.NewUpdatePersonService(personRepo)
	if err = updatePersonService.Execute(&person); err != nil {
		httpwr.Error(w, http.StatusInternalServerError, err)
		return
	}

	redis := cachedb.NewConnection()
	err = redis.Set(context.Background(), person.ID, person)
	if err != nil {
		log.Println(core.Red("Error writing search person ID to cache"), err)
	}

	httpwr.Response(w, http.StatusNoContent, nil)
}
