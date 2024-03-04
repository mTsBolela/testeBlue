package persons

type PersonRepository interface {
	Create(person *Person) (*Person, error)
	Update(person *Person) error
	Get(uid string) (*Person, error)
	GetAll(name string) (*[]Person, error)
	Delete(uid string) error
}
