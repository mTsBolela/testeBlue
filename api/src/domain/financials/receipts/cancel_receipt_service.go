package receipts

import (
	"contablue/src/domain/financials"
	"time"
)

type CancelReceiptService struct {
	Repo financials.FinancialRepository
}

func NewCancelReceiptService(repo financials.FinancialRepository) *CancelReceiptService {
	return &CancelReceiptService{repo}
}

func (service *CancelReceiptService) Execute(financial *financials.Financial) error {
	financial.UpdatedAt = time.Now().UTC()
	financial.Status = "canceled"
	err := service.Repo.Cancel(financial)
	if err != nil {
		return err
	}
	return nil
}
