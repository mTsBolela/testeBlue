package categories

import "github.com/joaocprofile/goh/core/errs"

type UpdateCategoryService struct {
	Repo CategoryRepository
}

func NewUpdateCategoryService(repo CategoryRepository) *UpdateCategoryService {
	return &UpdateCategoryService{repo}
}

func (service *UpdateCategoryService) Execute(category *Category) *errs.Error {
	if category.Type != "credit" && category.Type != "debit" {
		return errs.New("Invalid category type").BussinesError()
	}

	err := service.Repo.Update(category)
	if err != nil {
		return errs.New(err.Error()).SystemError()
	}
	return nil
}
