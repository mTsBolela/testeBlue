package receipts

import (
	"contablue/src/domain/financials"
	"contablue/src/infra/repositories"
	"context"
	"log"
	"net/http"

	"github.com/joaocprofile/goh/core"
	"github.com/joaocprofile/goh/database/cachedb"
	"github.com/joaocprofile/goh/httpwr"
)

func GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := httpwr.Params("id", w, r)
	if err != nil {
		return
	}

	redis := cachedb.NewConnection()
	paymentCache, err := redis.Get(context.Background(), id)
	if err == nil {
		httpwr.JSON(w, http.StatusOK, paymentCache)
		return
	}

	financialRepo := repositories.NewFinancialRepository()
	payment, err := financialRepo.Get(id)
	if err != nil {
		httpwr.Error(w, http.StatusInternalServerError, err)
		return
	}
	err = redis.Set(context.Background(), id, payment)
	if err != nil {
		log.Println(core.Red("Error writing Search Receipt ID to cache"), err)
	}

	httpwr.Response(w, http.StatusOK, payment)
}

func GetByIDResume(w http.ResponseWriter, r *http.Request) {
	id, err := httpwr.Params("id", w, r)
	if err != nil {
		return
	}

	redis := cachedb.NewConnection()
	paymentCache, err := redis.Get(context.Background(), id+"/resume")
	if err == nil {
		httpwr.JSON(w, http.StatusOK, paymentCache)
		return
	}

	financialRepo := repositories.NewFinancialRepository()
	payment, err := financialRepo.GetResume(id)
	if err != nil {
		httpwr.Error(w, http.StatusInternalServerError, err)
		return
	}
	err = redis.Set(context.Background(), id+"/resume", payment)
	if err != nil {
		log.Println(core.Red("Error writing Search Receipt ID to cache"), err)
	}

	httpwr.Response(w, http.StatusOK, payment)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	var filterStruct financials.FinancialFilter
	if err := httpwr.QueryToStruct(r, &filterStruct); err != nil {
		httpwr.Error(w, http.StatusOK, err)
		return
	}

	filterStruct.Type = "receipt"
	paymentRepo := repositories.NewFinancialRepository()
	paymentList, err := paymentRepo.GetAll(filterStruct)
	if err != nil {
		httpwr.Error(w, http.StatusInternalServerError, err)
		return
	}

	httpwr.Response(w, http.StatusOK, paymentList)
}

func GetAllResume(w http.ResponseWriter, r *http.Request) {
	var filterStruct financials.FinancialFilter
	if err := httpwr.QueryToStruct(r, &filterStruct); err != nil {
		httpwr.Error(w, http.StatusOK, err)
		return
	}

	filterStruct.Type = "receipt"
	paymentRepo := repositories.NewFinancialRepository()
	paymentList, err := paymentRepo.GetAllResume(filterStruct)
	if err != nil {
		httpwr.Error(w, http.StatusInternalServerError, err)
		return
	}

	httpwr.Response(w, http.StatusOK, paymentList)
}
