package receipts

import (
	"contablue/src/domain/financials"
	"contablue/src/domain/financials/receipts"
	"contablue/src/infra/repositories"
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/joaocprofile/goh/core"
	"github.com/joaocprofile/goh/database/cachedb"
	"github.com/joaocprofile/goh/httpwr"
)

func UpdateReceipt(w http.ResponseWriter, r *http.Request) {
	ID, err := httpwr.Params("id", w, r)
	if err != nil {
		return
	}

	var receipt financials.Financial
	if err = httpwr.ReadBody(w, r, &receipt); err != nil {
		return
	}
	if err = receipt.Prepare(); err != nil {
		httpwr.Error(w, http.StatusBadRequest, errors.New("Error Validate Receipt: "+err.Error()))
		return
	}
	receipt.ID = ID

	paymentRepo := repositories.NewFinancialRepository()
	personRepo := repositories.NewPersonRepository()
	categoryRepo := repositories.NewCategoryRepository()
	updateReceiptService := receipts.NewUpdateReceiptService(paymentRepo, personRepo, categoryRepo)
	if err := updateReceiptService.Execute(&receipt); err != nil {
		httpwr.Error(w, err.StatusCode, err.Err)
		return
	}

	redis := cachedb.NewConnection()
	if err = redis.Del(context.Background(), receipt.ID); err != nil {
		log.Println(core.Red("Error delete Receipt ID to cache"), err)
	}

	httpwr.Response(w, http.StatusNoContent, nil)
}
