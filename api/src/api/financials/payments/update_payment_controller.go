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

func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	ID, err := httpwr.Params("id", w, r)
	if err != nil {
		return
	}

	var payment financials.Financial
	if err = httpwr.ReadBody(w, r, &payment); err != nil {
		return
	}
	if err = payment.Prepare(); err != nil {
		httpwr.Error(w, http.StatusBadRequest, errors.New("Error Validate Payment: "+err.Error()))
		return
	}
	payment.ID = ID

	paymentRepo := repositories.NewFinancialRepository()
	personRepo := repositories.NewPersonRepository()
	categoryRepo := repositories.NewCategoryRepository()
	updatePaymentService := payments.NewUpdatePaymentService(paymentRepo, personRepo, categoryRepo)
	if err := updatePaymentService.Execute(&payment); err != nil {
		httpwr.Error(w, err.StatusCode, err.Err)
		return
	}

	redis := cachedb.NewConnection()
	if err = redis.Del(context.Background(), payment.ID); err != nil {
		log.Println(core.Red("Error delete Payment ID to cache"), err)
	}

	httpwr.Response(w, http.StatusNoContent, nil)
}
