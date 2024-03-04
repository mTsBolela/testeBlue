package categories

import (
	"time"

	"github.com/google/uuid"
	"github.com/joaocprofile/goh/core/errs"
)

type CreateCategoryService struct {
	Repo CategoryRepository
}

func NewCreateCategoryService(repo CategoryRepository) *CreateCategoryService {
	return &CreateCategoryService{repo}
}

func (service *CreateCategoryService) Execute(category *Category) (*Category, *errs.Error) {
	category.ID = uuid.New().String()
	category.Status = "active"
	category.CreatedAt = time.Now().UTC()

	newCategory, err := service.Repo.Create(category)
	if err != nil {
		return newCategory, errs.New(err.Error()).SystemError()
	}
	return newCategory, nil
}
