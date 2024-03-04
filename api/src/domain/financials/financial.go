package financials

import (
	"contablue/src/domain/categories"
	"contablue/src/domain/persons"
	"errors"
	"strings"
	"time"
)

type Financial struct {
	ID             string              `json:"id"`
	Type           string              `json:"type,omitempty"` // payment or receipt
	Category       categories.Category `json:"category"`
	Person         persons.Person      `json:"person"`
	Document       string              `json:"document"`
	DocumentValue  float64             `json:"document_value"`
	DischargeValue float64             `json:"discharge_value"`
	Description    string              `json:"description"`
	DueDate        time.Time           `json:"due_date"`
	DischargeDate  time.Time           `json:"discharge_date"`
	Status         string              `json:"status,omitempty"` // pending, paidout, canceled
	CreatedAt      time.Time           `json:"created_at,omitempty"`
	UpdatedAt      time.Time           `json:"updated_at,omitempty"`
}

func NewFinancial() *Financial {
	return &Financial{}
}

func (p Financial) Prepare() error {
	return p.validate()
}

func (p Financial) validate() error {
	if p.Category.ID == "" {
		return errors.New("Category Id is required")
	}
	if p.Person.ID == "" {
		return errors.New("Person Id Id is required")
	}
	if p.Document == "" {
		return errors.New("Document Id is required")
	}
	if p.Description == "" {
		return errors.New("Description Id is required")
	}
	if p.DocumentValue == 0 {
		return errors.New("Document Value is required")
	}

	p.formatFields()
	return nil
}

func (p *Financial) formatFields() {
	p.Description = strings.TrimSpace(p.Description)
	p.Document = strings.TrimSpace(p.Document)
}
