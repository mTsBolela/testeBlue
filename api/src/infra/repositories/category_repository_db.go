package repositories

import (
	"contablue/src/domain/categories"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/joaocprofile/goh/database/sqlg"
)

type CategoryRepositoryDB struct {
}

func NewCategoryRepository() *CategoryRepositoryDB {
	return &CategoryRepositoryDB{}
}

func (repository CategoryRepositoryDB) Create(category *categories.Category) (*categories.Category, error) {
	sql := "INSERT INTO categories (id, description, type, status) VALUES ($1, $2, $3, $4)"

	query := sqlg.NewQuery(context.Background())
	query.AddSQL(sql,
		category.ID,
		category.Description,
		category.Type,
		category.Status,
	)
	if err := query.Execute(); err != nil {
		return nil, errors.New("Error persisting a new Category " + err.Error())
	}

	defer query.Close()
	return category, nil
}

func (repository CategoryRepositoryDB) Get(uuid string) (*categories.Category, error) {
	sql := "SELECT id, description, type, status, created_at FROM categories WHERE id = $1"

	query := sqlg.NewQuery(context.Background())
	query.AddSQL(sql, uuid)
	row, err := query.Open()
	if err != nil {
		return nil, errors.New("Error finding category by ID: " + err.Error())
	}

	var category categories.Category
	if row.Next() {
		if err := row.Scan(
			&category.ID,
			&category.Description,
			&category.Type,
			&category.Status,
			&category.CreatedAt); err != nil {
			return nil, errors.New("Error finding category by ID: " + err.Error())
		}
	}

	defer query.Close()
	defer row.Close()
	return &category, nil
}

func (repository CategoryRepositoryDB) GetAll(description string) (*[]categories.Category, error) {
	filter := fmt.Sprintf("%%%s%%", description)
	sql := "SELECT id, description, type, status, created_at FROM categories WHERE LOWER(description) LIKE $1"

	categoriesList := []categories.Category{}
	query := sqlg.NewQuery(context.Background())
	query.AddSQL(sql, strings.ToLower(filter))
	rows, err := query.Open()
	if err != nil {
		return &categoriesList, errors.New("Error listing Description: " + err.Error())
	}

	for rows.Next() {
		var category categories.Category
		if err := rows.Scan(
			&category.ID,
			&category.Description,
			&category.Type,
			&category.Status,
			&category.CreatedAt); err != nil {
			return &categoriesList, errors.New("Error listing Descriptions: " + err.Error())
		}
		categoriesList = append(categoriesList, category)
	}

	defer rows.Close()
	defer query.Close()
	return &categoriesList, nil
}

func (repository CategoryRepositoryDB) Update(category *categories.Category) error {
	sql := "UPDATE categories SET description = $1, type = $2, status = $3 WHERE id = $4"

	query := sqlg.NewQuery(context.Background())
	query.AddSQL(sql, category.Description, category.Type, category.Status, category.ID)
	if err := query.Execute(); err != nil {
		return errors.New("Error updating Category " + err.Error())
	}

	defer query.Close()
	return nil
}

func (repository CategoryRepositoryDB) Delete(uuid string) error {
	sql := "DELETE FROM categories WHERE id = $1"

	query := sqlg.NewQuery(context.Background())
	query.AddSQL(sql, uuid)
	if err := query.Execute(); err != nil {
		return errors.New("Error Deleting Category (Statement) " + err.Error())
	}

	defer query.Close()
	return nil
}
