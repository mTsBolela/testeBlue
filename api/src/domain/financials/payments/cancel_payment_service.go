package payments

import (
	"contablue/src/domain/financials"
	"time"
)

type CancelPaymentService struct {
	Repo financials.FinancialRepository
}

func NewCancelPaymentService(repo financials.FinancialRepository) *CancelPaymentService {
	return &CancelPaymentService{repo}
}

func (service *CancelPaymentService) Execute(financial *financials.Financial) error {
	financial.UpdatedAt = time.Now().UTC()
	financial.Status = "canceled"
	err := service.Repo.Cancel(financial)
	if err != nil {
		return err
	}
	return nil
}
