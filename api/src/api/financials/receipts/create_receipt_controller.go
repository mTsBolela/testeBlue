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

func CreateReceipt(w http.ResponseWriter, r *http.Request) {
	var financial financials.Financial
	if err := httpwr.ReadBody(w, r, &financial); err != nil {
		return
	}

	if err := financial.Prepare(); err != nil {
		httpwr.Error(w, http.StatusBadRequest, errors.New("Error Validate Financial Register: "+err.Error()))
		return
	}

	receiptRepo := repositories.NewFinancialRepository()
	personRepo := repositories.NewPersonRepository()
	categoryRepo := repositories.NewCategoryRepository()
	createReceiptService := receipts.NewCreateReceiptService(receiptRepo, personRepo, categoryRepo)
	createdReceipt, err := createReceiptService.Execute(&financial)
	if err != nil {
		httpwr.Error(w, err.StatusCode, err.Err)
		return
	}

	redis := cachedb.NewConnection()
	if err := redis.Set(context.Background(), createdReceipt.ID, createdReceipt); err != nil {
		log.Println(core.Red("Error writing search Receipt ID to cache"), err)
	}

	httpwr.Response(w, http.StatusCreated, createdReceipt)
}
