package persons

type UpdatePersonService struct {
	Repo PersonRepository
}

func NewUpdatePersonService(repo PersonRepository) *UpdatePersonService {
	return &UpdatePersonService{repo}
}

func (service *UpdatePersonService) Execute(person *Person) error {
	if err := service.Repo.Update(person); err != nil {
		return err
	}
	return nil
}
