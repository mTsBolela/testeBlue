package persons

import (
	"time"

	"github.com/google/uuid"
)

type CreatePersonService struct {
	Repo PersonRepository
}

func NewCreatePersonService(repo PersonRepository) *CreatePersonService {
	return &CreatePersonService{repo}
}

func (service *CreatePersonService) Execute(person *Person) (*Person, error) {
	person.ID = uuid.New().String()
	person.Status = "active"
	person.CreatedAt = time.Now().UTC()

	newCategory, err := service.Repo.Create(person)
	if err != nil {
		return newCategory, err
	}
	return newCategory, nil
}
