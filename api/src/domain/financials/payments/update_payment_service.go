package payments

import (
	"contablue/src/domain/categories"
	"contablue/src/domain/financials"
	"contablue/src/domain/persons"
	"time"

	"github.com/joaocprofile/goh/core/errs"
)

type UpdatePaymentService struct {
	PaymentRepo  financials.FinancialRepository
	PersonRepo   persons.PersonRepository
	CategoryRepo categories.CategoryRepository
}

func NewUpdatePaymentService(
	paymentRepo financials.FinancialRepository,
	personRepo persons.PersonRepository,
	categoryRepo categories.CategoryRepository) *UpdatePaymentService {
	return &UpdatePaymentService{
		PaymentRepo:  paymentRepo,
		PersonRepo:   personRepo,
		CategoryRepo: categoryRepo,
	}
}

func (service *UpdatePaymentService) Execute(financial *financials.Financial) *errs.Error {
	category, _ := service.CategoryRepo.Get(financial.Category.ID)
	if category.Status == "inactive" {
		return errs.New("Category is inactive").BussinesError()
	}
	if category.Type == "credit" {
		return errs.New("Invalid type to this operation").BussinesError()
	}

	person, _ := service.PersonRepo.Get(financial.Person.ID)
	if person.Status == "inactive" {
		return errs.New("Person is inactive").BussinesError()
	}
	paymentDB, _ := service.PaymentRepo.Get(financial.ID)
	if paymentDB.Status == "paidout" || paymentDB.Status == "canceled" {
		return errs.New("Payment already paid or canceled").BussinesError()
	}

	financial.UpdatedAt = time.Now().UTC()
	err := service.PaymentRepo.Update(financial)
	if err != nil {
		return errs.New(err.Error()).SystemError()
	}
	return nil
}
