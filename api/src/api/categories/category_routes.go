package categories

import (
	"net/http"

	"github.com/joaocprofile/goh/core"
)

var Routes = []core.Route{
	{
		URI:            "/categories",
		Method:         http.MethodPost,
		Function:       CreateCategory,
		Authentication: false,
	},
	{
		URI:            "/categories",
		Method:         http.MethodGet,
		Function:       GetAll,
		Authentication: false,
	},
	{
		URI:            "/categories/{id}",
		Method:         http.MethodGet,
		Function:       GetByID,
		Authentication: false,
	},
	{
		URI:            "/categories/{id}",
		Method:         http.MethodPut,
		Function:       UpdateCategory,
		Authentication: false,
	},
	{
		URI:            "/categories/{id}",
		Method:         http.MethodDelete,
		Function:       DeleteCategory,
		Authentication: false,
	},
}
