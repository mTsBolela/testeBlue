package categories

import (
	"strings"
	"time"

	"github.com/joaocprofile/goh/core/errs"
)

type Category struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Type        string    `json:"type"`   // credit, debit
	Status      string    `json:"status"` // Active / Inactive
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

func NewCategory() *Category {
	return &Category{}
}

func (c Category) Prepare() *errs.Error {
	return c.validate()
}

func (c Category) validate() *errs.Error {
	if c.Description == "" {
		return errs.New("Description is required").BussinesError()
	}
	if c.Type == "" {
		return errs.New("Type is required").BussinesError()
	}
	if c.Type != "credit" && c.Type != "debit" {
		return errs.New("Invalid category type").BussinesError()
	}
	if c.Status != "active" && c.Status != "inactive" {
		return errs.New("Invalid category Status").BussinesError()
	}

	c.Description = strings.TrimSpace(c.Description)
	return nil
}
