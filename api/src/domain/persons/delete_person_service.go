package persons

type DeletePersonService struct {
	Repo PersonRepository
}

func NewDeletePersonService(repo PersonRepository) *DeletePersonService {
	return &DeletePersonService{repo}
}

func (service *DeletePersonService) Execute(uuid string) error {
	if err := service.Repo.Delete(uuid); err != nil {
		return err
	}
	return nil
}
