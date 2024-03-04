package receipts

import (
	"contablue/src/domain/categories"
	"contablue/src/domain/financials"
	"contablue/src/domain/persons"
	"time"

	"github.com/google/uuid"
	"github.com/joaocprofile/goh/core/errs"
)

type CreateReceiptService struct {
	ReceiptRepo  financials.FinancialRepository
	PersonRepo   persons.PersonRepository
	CategoryRepo categories.CategoryRepository
}

func NewCreateReceiptService(
	receiptRepo financials.FinancialRepository,
	personRepo persons.PersonRepository,
	categoryRepo categories.CategoryRepository) *CreateReceiptService {
	return &CreateReceiptService{
		ReceiptRepo:  receiptRepo,
		PersonRepo:   personRepo,
		CategoryRepo: categoryRepo,
	}
}

func (service *CreateReceiptService) Execute(financial *financials.Financial) (*financials.Financial, *errs.Error) {
	category, _ := service.CategoryRepo.Get(financial.Category.ID)
	if category.ID == "" {
		return nil, errs.New("Invalid Category").BussinesError()
	}
	if category.Status == "inactive" {
		return nil, errs.New("Category is inactive").BussinesError()
	}
	if category.Type == "debit" {
		return nil, errs.New("Invalid type to this operation").BussinesError()
	}

	person, _ := service.PersonRepo.Get(financial.Person.ID)
	if person.ID == "" {
		return nil, errs.New("Invalid Person").BussinesError()
	}
	if person.Status == "inactive" {
		return nil, errs.New("Can't continue, Person is inactive").BussinesError()
	}

	financial.ID = uuid.New().String()
	financial.Type = "receipt"
	financial.Status = "pending"
	financial.CreatedAt = time.Now().UTC()

	newPayment, err := service.ReceiptRepo.Create(financial)
	if err != nil {
		return newPayment, errs.New(err.Error()).SystemError()
	}
	return newPayment, nil
}
