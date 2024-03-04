package payments

import (
	"contablue/src/domain/financials"
	"contablue/src/domain/financials/payments"
	"contablue/src/infra/repositories"
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/joaocprofile/goh/core"
	"github.com/joaocprofile/goh/database/cachedb"
	"github.com/joaocprofile/goh/httpwr"
)

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var financial financials.Financial
	if err := httpwr.ReadBody(w, r, &financial); err != nil {
		return
	}

	if err := financial.Prepare(); err != nil {
		httpwr.Error(w, http.StatusBadRequest, errors.New("Error Validate Financial Register: "+err.Error()))
		return
	}

	paymentRepo := repositories.NewFinancialRepository()
	personRepo := repositories.NewPersonRepository()
	categoryRepo := repositories.NewCategoryRepository()

	createPaymentService := payments.NewCreatePaymentService(paymentRepo, personRepo, categoryRepo)
	createdPayment, err := createPaymentService.Execute(&financial)
	if err != nil {
		httpwr.Error(w, err.StatusCode, err.Err)
		return
	}

	redis := cachedb.NewConnection()
	if err := redis.Set(context.Background(), createdPayment.ID, createdPayment); err != nil {
		log.Println(core.Red("Error writing search Payment ID to cache"), err)
	}

	httpwr.Response(w, http.StatusCreated, createdPayment)
}
