package receipts

import (
	"net/http"

	"github.com/joaocprofile/goh/core"
)

var Routes = []core.Route{
	{
		URI:            "/receipts",
		Method:         http.MethodGet,
		Function:       GetAll,
		Authentication: false,
	},
	{
		URI:            "/receipts/resume",
		Method:         http.MethodGet,
		Function:       GetAllResume,
		Authentication: false,
	},
	{
		URI:            "/receipts/{id}",
		Method:         http.MethodGet,
		Function:       GetByID,
		Authentication: false,
	},
	{
		URI:            "/receipts/{id}/resume",
		Method:         http.MethodGet,
		Function:       GetByIDResume,
		Authentication: false,
	},

	{
		URI:            "/receipts",
		Method:         http.MethodPost,
		Function:       CreateReceipt,
		Authentication: false,
	},
	{
		URI:            "/receipts/{id}",
		Method:         http.MethodPut,
		Function:       UpdateReceipt,
		Authentication: false,
	},
	{
		URI:            "/receipts/{id}/cancel",
		Method:         http.MethodPost,
		Function:       CancelReceipt,
		Authentication: false,
	},
	{
		URI:            "/receipts/{id}/discharge",
		Method:         http.MethodPost,
		Function:       CancelReceipt,
		Authentication: false,
	},
}
