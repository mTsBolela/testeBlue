package persons

import (
	"net/http"

	"github.com/joaocprofile/goh/core"
)

var Routes = []core.Route{
	{
		URI:            "/persons/{id}",
		Method:         http.MethodGet,
		Function:       GetByID,
		Authentication: false,
	},
	{
		URI:            "/persons",
		Method:         http.MethodGet,
		Function:       GetAll,
		Authentication: false,
	},
	{
		URI:            "/persons",
		Method:         http.MethodPost,
		Function:       CreatePerson,
		Authentication: false,
	},
	{
		URI:            "/persons/{id}",
		Method:         http.MethodPut,
		Function:       UpdatePerson,
		Authentication: false,
	},
	{
		URI:            "/persons/{id}",
		Method:         http.MethodDelete,
		Function:       DeletePerson,
		Authentication: false,
	},
}
