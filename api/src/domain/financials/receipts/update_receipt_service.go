package receipts

import (
	"contablue/src/domain/categories"
	"contablue/src/domain/financials"
	"contablue/src/domain/persons"
	"time"

	"github.com/joaocprofile/goh/core/errs"
)

type UpdateReceiptService struct {
	ReceiptRepo  financials.FinancialRepository
	PersonRepo   persons.PersonRepository
	CategoryRepo categories.CategoryRepository
}

func NewUpdateReceiptService(
	receiptRepo financials.FinancialRepository,
	personRepo persons.PersonRepository,
	categoryRepo categories.CategoryRepository) *UpdateReceiptService {
	return &UpdateReceiptService{
		ReceiptRepo:  receiptRepo,
		PersonRepo:   personRepo,
		CategoryRepo: categoryRepo,
	}
}

func (service *UpdateReceiptService) Execute(financial *financials.Financial) *errs.Error {
	category, _ := service.CategoryRepo.Get(financial.Category.ID)
	if category.Status == "inactive" {
		return errs.New("Category is inactive").BussinesError()
	}
	if category.Type == "debit" {
		return errs.New("Invalid type to this operation").BussinesError()
	}

	person, _ := service.PersonRepo.Get(financial.Person.ID)
	if person.Status == "inactive" {
		return errs.New("Person is inactive").BussinesError()
	}
	receiptDB, _ := service.ReceiptRepo.Get(financial.ID)
	if receiptDB.Status == "paidout" || receiptDB.Status == "canceled" {
		return errs.New("Receipt already paid or canceled").BussinesError()
	}

	financial.UpdatedAt = time.Now().UTC()
	err := service.ReceiptRepo.Update(financial)
	if err != nil {
		return errs.New(err.Error()).SystemError()
	}
	return nil
}
