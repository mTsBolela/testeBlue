package main

import (
	"contablue/src/api/categories"
	"contablue/src/api/financials/payments"
	"contablue/src/api/financials/receipts"
	"contablue/src/api/persons"

	"github.com/joaocprofile/goh/environment"
	"github.com/joaocprofile/goh/httprest"
)

func main() {
	environment.Inicialize()

	apiRest := httprest.NewHttpRest().AddDatabase().AddCache()
	apiRest.AddRoutes(categories.Routes)
	apiRest.AddRoutes(persons.Routes)
	apiRest.AddRoutes(payments.Routes)
	apiRest.AddRoutes(receipts.Routes)
	apiRest.ListenAndServe()
}
