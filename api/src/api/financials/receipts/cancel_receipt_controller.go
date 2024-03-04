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

func CancelReceipt(w http.ResponseWriter, r *http.Request) {
	ID, err := httpwr.Params("id", w, r)
	if err != nil {
		return
	}

	receipt := financials.Financial{}
	receipt.ID = ID

	financialRepo := repositories.NewFinancialRepository()
	cancelReceiptService := payments.NewCancelPaymentService(financialRepo)
	err = cancelReceiptService.Execute(&receipt)
	if err != nil {
		httpwr.Error(w, http.StatusInternalServerError, err)
		return
	}

	redis := cachedb.NewConnection()
	if err = redis.Del(context.Background(), receipt.ID); err != nil {
		log.Println(core.Red("Error delete Receipt ID to cache"), err)
	}

	httpwr.Response(w, http.StatusNoContent, nil)
}
