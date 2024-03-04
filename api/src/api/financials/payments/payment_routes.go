package payments

import (
	"net/http"

	"github.com/joaocprofile/goh/core"
)

var Routes = []core.Route{
	{
		URI:            "/payments",
		Method:         http.MethodGet,
		Function:       GetAll,
		Authentication: false,
	},
	{
		URI:            "/payments/resume",
		Method:         http.MethodGet,
		Function:       GetAllResume,
		Authentication: false,
	},
	{
		URI:            "/payments/{id}",
		Method:         http.MethodGet,
		Function:       GetByID,
		Authentication: false,
	},
	{
		URI:            "/payments/{id}/resume",
		Method:         http.MethodGet,
		Function:       GetByIDResume,
		Authentication: false,
	},

	{
		URI:            "/payments",
		Method:         http.MethodPost,
		Function:       CreatePayment,
		Authentication: false,
	},
	{
		URI:            "/payments/{id}",
		Method:         http.MethodPut,
		Function:       UpdatePayment,
		Authentication: false,
	},
	{
		URI:            "/payments/{id}/cancel",
		Method:         http.MethodPost,
		Function:       CancelPayment,
		Authentication: false,
	},
	{
		URI:            "/payments/{id}/discharge",
		Method:         http.MethodPost,
		Function:       DischargePayment,
		Authentication: false,
	},
}
