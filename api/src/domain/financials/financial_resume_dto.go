package financials

import (
	"time"
)

type PersonResume struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CategoryResume struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type FinancialResume struct {
	ID             string         `json:"id"`
	Type           string         `json:"type"` // payment or receipt
	Category       CategoryResume `json:"category"`
	Person         PersonResume   `json:"person"`
	Document       string         `json:"document"`
	DocumentValue  float64        `json:"document_value"`
	DischargeValue float64        `json:"discharge_value"`
	Description    string         `json:"description"`
	DueDate        time.Time      `json:"due_date"`
	DischargeDate  time.Time      `json:"discharge_date"`
	Status         string         `json:"status,omitempty"` // pending, paidout, canceled
	CreatedAt      time.Time      `json:"created_at,omitempty"`
	UpdatedAt      time.Time      `json:"updated_at,omitempty"`
}
