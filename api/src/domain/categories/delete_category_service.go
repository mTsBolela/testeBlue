package categories

type DeleteCategoryService struct {
	Repo CategoryRepository
}

func NewDeleteCategoryService(repo CategoryRepository) *DeleteCategoryService {
	return &DeleteCategoryService{repo}
}

func (service *DeleteCategoryService) Execute(uuid string) error {
	err := service.Repo.Delete(uuid)
	if err != nil {
		return err
	}
	return nil
}
