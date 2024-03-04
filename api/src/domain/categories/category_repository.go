package categories

type CategoryRepository interface {
	Create(category *Category) (*Category, error)
	Update(category *Category) error
	Get(uid string) (*Category, error)
	GetAll(description string) (*[]Category, error)
	Delete(uid string) error
}
