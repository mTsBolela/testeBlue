package persons

import (
	"contablue/src/domain/persons"
	"contablue/src/infra/repositories"
	"net/http"

	"github.com/joaocprofile/goh/httpwr"
)

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	ID, err := httpwr.Params("id", w, r)
	if err != nil {
		return
	}

	personRepo := repositories.NewPersonRepository()
	deletePersonService := persons.NewDeletePersonService(personRepo)
	if err = deletePersonService.Execute(ID); err != nil {
		httpwr.Error(w, http.StatusInternalServerError, err)
		return
	}

	httpwr.Response(w, http.StatusNoContent, nil)
}
