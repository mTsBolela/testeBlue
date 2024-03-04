package persons

import (
	"strings"
	"time"

	"github.com/joaocprofile/goh/core/errs"
)

type Person struct {
	ID         string    `json:"id"`
	Customer   string    `json:"customer"`
	Provider   string    `json:"provider"`
	Document   string    `json:"document"`
	Name       string    `json:"name"`
	Cep        string    `json:"cep"`
	Address    string    `json:"address"`
	State      string    `json:"state"`
	City       string    `json:"city"`
	Complement string    `json:"complement"`
	Status     string    `json:"status"` // Active / Inactive
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

func NewPerson() *Person {
	return &Person{}
}

func (p Person) Prepare() *errs.Error {
	return p.validate()
}

func (p Person) validate() *errs.Error {
	if p.Name == "" {
		return errs.New("Name is required").BussinesError()
	}
	if p.Customer != "true" && p.Customer != "false" {
		return errs.New("Customer is required").BussinesError()
	}
	if p.Provider != "true" && p.Provider != "false" {
		return errs.New("Provider is required").BussinesError()
	}
	if p.Status != "active" && p.Status != "inactive" {
		return errs.New("Invalid person status").BussinesError()
	}
	p.formatFields()
	return nil
}

func (p *Person) formatFields() {
	p.Name = strings.TrimSpace(p.Name)
	p.Address = strings.TrimSpace(p.Address)
	p.State = strings.TrimSpace(p.State)
	p.City = strings.TrimSpace(p.City)
	p.Complement = strings.TrimSpace(p.Complement)
}
