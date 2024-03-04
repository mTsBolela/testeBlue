package repositories

import (
	"contablue/src/domain/persons"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/joaocprofile/goh/database/sqlg"
)

type PersonRepositoryDB struct{}

func NewPersonRepository() *PersonRepositoryDB {
	return &PersonRepositoryDB{}
}

func (repository PersonRepositoryDB) Create(person *persons.Person) (*persons.Person, error) {
	sql := `INSERT INTO persons 
	            (id, name, customer, provider, document, cep, address, state, city, complement, status, created_at) 
						VALUES 
						  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	query := sqlg.NewQuery(context.Background())
	query.AddSQL(sql,
		person.ID,
		person.Name,
		person.Customer,
		person.Provider,
		person.Document,
		person.Cep,
		person.Address,
		person.State,
		person.City,
		person.Complement,
		person.Status,
		person.CreatedAt,
	)
	if err := query.Execute(); err != nil {
		return nil, errors.New("Error persisting a new Person " + err.Error())
	}

	defer query.Close()
	return person, nil
}

func (repository PersonRepositoryDB) Get(uuid string) (*persons.Person, error) {
	sql := ` SELECT 
	           id, name, customer, provider, document, cep, address, state, city, complement, status, created_at 
					 FROM 
					   persons 
					 WHERE 
					   id = $1
	`
	query := sqlg.NewQuery(context.Background())
	query.AddSQL(sql, uuid)
	row, err := query.Open()
	if err != nil {
		return nil, errors.New("Error finding Person by ID: " + err.Error())
	}

	var person persons.Person
	if row.Next() {
		if err := row.Scan(
			&person.ID,
			&person.Name,
			&person.Customer,
			&person.Provider,
			&person.Document,
			&person.Cep,
			&person.Address,
			&person.State,
			&person.City,
			&person.Complement,
			&person.Status,
			&person.CreatedAt,
		); err != nil {
			return nil, errors.New("Error finding Person by ID: " + err.Error())
		}
	}

	defer query.Close()
	defer row.Close()
	return &person, nil
}

func (repository PersonRepositoryDB) GetAll(name string) (*[]persons.Person, error) {
	filterName := fmt.Sprintf("%%%s%%", name)
	sql := ` SELECT 
	           id, name, customer, provider, document, cep, address, state, city, complement, status, created_at 
					 FROM 
					   persons 
					 WHERE 
					   LOWER(name) LIKE $1
	`
	personList := []persons.Person{}
	query := sqlg.NewQuery(context.Background())
	query.AddSQL(sql, strings.ToLower(filterName))
	rows, err := query.Open()
	if err != nil {
		return &personList, errors.New("Error listing Name: " + err.Error())
	}

	for rows.Next() {
		var person persons.Person
		if err := rows.Scan(
			&person.ID,
			&person.Name,
			&person.Customer,
			&person.Provider,
			&person.Document,
			&person.Cep,
			&person.Address,
			&person.State,
			&person.City,
			&person.Complement,
			&person.Status,
			&person.CreatedAt,
		); err != nil {
			return &personList, errors.New("Error listing Descriptions: " + err.Error())
		}
		personList = append(personList, person)
	}

	defer query.Close()
	defer rows.Close()
	return &personList, nil
}

func (repository PersonRepositoryDB) Update(person *persons.Person) error {
	sql := `UPDATE persons SET 
	            name = $1, 
							customer = $2, 
							provider = $3,
							document = $4, 
							cep = $5, 
							address = $6, 
							state = $7,
							city = $8,
							complement = $9,
							status = $10 
						WHERE 
						  ID = $11
	`
	query := sqlg.NewQuery(context.Background())
	query.AddSQL(sql,
		person.Name,
		person.Customer,
		person.Provider,
		person.Document,
		person.Cep,
		person.Address,
		person.State,
		person.City,
		person.Complement,
		person.Status,
		person.ID,
	)
	if err := query.Execute(); err != nil {
		return errors.New("Error updating Person " + err.Error())
	}

	defer query.Close()
	return nil
}

func (repository PersonRepositoryDB) Delete(uuid string) error {
	sql := "DELETE FROM persons WHERE id = $1"

	query := sqlg.NewQuery(context.Background())
	query.AddSQL(sql, uuid)
	if err := query.Execute(); err != nil {
		return errors.New("Error Deleting Person " + err.Error())
	}

	defer query.Close()
	return nil
}
