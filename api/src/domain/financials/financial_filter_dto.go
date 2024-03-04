package financials

type FinancialFilter struct {
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	InitialDate string `json:"initialDate,omitempty"`
	FinalDate   string `json:"finalDate,omitempty"`
}
