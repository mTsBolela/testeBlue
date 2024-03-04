package payments

import (
	"contablue/src/domain/financials"
	"time"

	"github.com/joaocprofile/goh/core/errs"
)

type DischargePaymentService struct {
	Repo financials.FinancialRepository
}

func NewDischargePaymentService(repo financials.FinancialRepository) *DischargePaymentService {
	return &DischargePaymentService{repo}
}

func (service *DischargePaymentService) Execute(financial *financials.Financial) *errs.Error {

	if financial.DischargeDate.Before(financial.DueDate) {
		return errs.New("The Dischanrge date cannot be less than the Due Date").BussinesError()
	}
	if financial.DischargeValue == 0 {
		return errs.New("Discharge Value is required").BussinesError()
	}

	financial.UpdatedAt = time.Now().UTC()
	financial.Status = "paidout"

	err := service.Repo.Update(financial)
	if err != nil {
		return errs.New(err.Error()).SystemError()
	}
	return nil
}
