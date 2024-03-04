package financials

type FinancialRepository interface {
	Create(financial *Financial) (*Financial, error)
	Update(financial *Financial) error
	Get(uid string) (*Financial, error)
	GetResume(uid string) (*FinancialResume, error)
	GetAll(filter FinancialFilter) (*[]Financial, error)
	GetAllResume(filter FinancialFilter) (*[]FinancialResume, error)
	Cancel(financial *Financial) error
	Discharge(financial *Financial) error
}
