package payments

import (
	"contablue/src/domain/financials"
	"contablue/src/domain/financials/payments"
	"contablue/src/infra/repositories"
	"context"
	"log"
	"net/http"

	"github.com/joaocprofile/goh/core"
	"github.com/joaocprofile/goh/database/cachedb"
	"github.com/joaocprofile/goh/httpwr"
)

func CancelPayment(w http.ResponseWriter, r *http.Request) {
	ID, err := httpwr.Params("id", w, r)
	if err != nil {
		return
	}

	payment := financials.Financial{}
	payment.ID = ID

	financialRepo := repositories.NewFinancialRepository()
	cancelPaymentService := payments.NewCancelPaymentService(financialRepo)
	err = cancelPaymentService.Execute(&payment)
	if err != nil {
		httpwr.Error(w, http.StatusInternalServerError, err)
		return
	}

	redis := cachedb.NewConnection()
	if err = redis.Del(context.Background(), payment.ID); err != nil {
		log.Println(core.Red("Error delete Payment ID to cache"), err)
	}

	httpwr.Response(w, http.StatusNoContent, nil)
}
