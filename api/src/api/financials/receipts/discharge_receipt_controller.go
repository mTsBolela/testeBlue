package receipts

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

func DischargePayment(w http.ResponseWriter, r *http.Request) {
	ID, err := httpwr.Params("id", w, r)
	if err != nil {
		return
	}

	var payment financials.Financial
	if err = httpwr.ReadBody(w, r, &payment); err != nil {
		return
	}
	payment.ID = ID

	paymentRepo := repositories.NewFinancialRepository()
	dischargePaymentService := payments.NewDischargePaymentService(paymentRepo)
	if err := dischargePaymentService.Execute(&payment); err != nil {
		httpwr.Error(w, err.StatusCode, err.Err)
		return
	}

	redis := cachedb.NewConnection()
	if err = redis.Del(context.Background(), payment.ID); err != nil {
		log.Println(core.Red("Error delete Payment ID to cache"), err)
	}

	httpwr.Response(w, http.StatusNoContent, nil)
}
