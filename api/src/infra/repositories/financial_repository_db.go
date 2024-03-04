package repositories

import (
	"contablue/src/domain/financials"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/joaocprofile/goh/database/sqlg"
)

type FinancialRepositoryDB struct{}

func NewFinancialRepository() *FinancialRepositoryDB {
	return &FinancialRepositoryDB{}
}

func (repository FinancialRepositoryDB) Create(financial *financials.Financial) (*financials.Financial, error) {
	sql := `INSERT INTO financials 
	            ( id, 
								type, 
								category_id, 
								person_id, 
								document, 
								document_value, 
								discharge_value, 
								description, 
								due_date, 
								discharge_date, 
								status, 
								created_at) 
						VALUES 
						  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
	query := sqlg.NewQuery(context.Background())
	query.AddSQL(sql,
		financial.ID,
		financial.Type,
		financial.Category.ID,
		financial.Person.ID,
		financial.Document,
		financial.DocumentValue,
		financial.DischargeValue,
		financial.Description,
		financial.DueDate,
		financial.DischargeDate,
		financial.Status,
		financial.CreatedAt,
	)
	if err := query.Execute(); err != nil {
		return nil, errors.New("Error persisting a new " + financial.Type + " " + err.Error())
	}

	defer query.Close()
	return financial, nil
}

func (repository FinancialRepositoryDB) Get(uuid string) (*financials.Financial, error) {
	sql := ` SELECT 
	           f.id, 
						 f.type, 
						 F.category_id, 
						 c.description,
						 c.type,
						 c.created_at,
						 f.person_id, 
						 p.name,						 
						 p.customer, 
						 p.provider, 
						 p.document, 
						 p.cep, 
						 p.address, 
						 p.state, 
						 p.city, 
						 p.complement,
						 p.created_at,
						 f.document, 
						 f.document_value,
						 f.discharge_value,
						 f.description, 
						 f.due_date, 
						 f.discharge_date, 
						 f.status, 
						 f.created_at
					 FROM 
					   financials f
					INNER JOIN 
						categories c ON c.id = f.category_id
					INNER JOIN 
						persons p ON p.id = f.person_id
					 WHERE 
					   f.id = $1
	`
	query := sqlg.NewQuery(context.Background())
	query.AddSQL(sql, uuid)
	row, err := query.Open()
	if err != nil {
		return nil, errors.New("Error finding Financial register by ID: " + err.Error())
	}
	var financial financials.Financial
	if row.Next() {
		if err := row.Scan(
			&financial.ID,
			&financial.Type,
			&financial.Category.ID,
			&financial.Category.Type,
			&financial.Category.Description,
			&financial.Category.CreatedAt,
			&financial.Person.ID,
			&financial.Person.Name,
			&financial.Person.Customer,
			&financial.Person.Provider,
			&financial.Person.Document,
			&financial.Person.Cep,
			&financial.Person.Address,
			&financial.Person.State,
			&financial.Person.City,
			&financial.Person.Complement,
			&financial.Person.CreatedAt,
			&financial.Document,
			&financial.DocumentValue,
			&financial.DischargeValue,
			&financial.Description,
			&financial.DueDate,
			&financial.DischargeDate,
			&financial.Status,
			&financial.CreatedAt,
		); err != nil {
			return nil, errors.New("Error finding Financial register by ID: " + err.Error())
		}
	}

	defer query.Close()
	defer row.Close()
	return &financial, nil
}

func (repository FinancialRepositoryDB) GetResume(uuid string) (*financials.FinancialResume, error) {
	sql := ` SELECT 
	           f.id, 
						 f.type, 
						 f.category_id, 
						 c.description,						 
						 f.person_id, 
						 p.name,						 						 
						 f.document, 
						 f.document_value,
						 f.discharge_value,
						 f.description, 
						 f.due_date, 
						 f.discharge_date, 
						 f.status, 
						 f.created_at
					 FROM 
					   financials f
					INNER JOIN 
						categories c ON c.id = f.category_id
					INNER JOIN 
						persons p ON p.id = f.person_id
					 WHERE 
					   f.id = $1
	`
	query := sqlg.NewQuery(context.Background())
	query.AddSQL(sql, uuid)
	row, err := query.Open()
	if err != nil {
		return nil, errors.New("Error finding Financial register by ID: " + err.Error())
	}
	var financial financials.FinancialResume
	if row.Next() {
		if err := row.Scan(
			&financial.ID,
			&financial.Type,
			&financial.Category.Id,
			&financial.Category.Description,
			&financial.Person.Id,
			&financial.Person.Name,
			&financial.Document,
			&financial.DocumentValue,
			&financial.DischargeValue,
			&financial.Description,
			&financial.DueDate,
			&financial.DischargeDate,
			&financial.Status,
			&financial.CreatedAt,
		); err != nil {
			return nil, errors.New("Error finding Financial register by ID: " + err.Error())
		}
	}

	defer query.Close()
	defer row.Close()
	return &financial, nil
}

func (repository FinancialRepositoryDB) GetAll(filter financials.FinancialFilter) (*[]financials.Financial, error) {
	filterType := filter.Type
	description := fmt.Sprintf("%%%s%%", filter.Description)

	sql := ` SELECT 
				  	 f.id, 
						 f.type, 
						 f.category_id, 
						 c.description,
						 c.type,
						 c.created_at,
						 f.person_id, 
						 p.name,						 
						 p.customer, 
						 p.provider, 
						 p.document, 
						 p.cep, 
						 p.address, 
						 p.state, 
						 p.city, 
						 p.complement,
						 p.created_at,
						 f.document, 
						 f.document_value,
						 f.discharge_value,
						 f.description, 
						 f.due_date, 
						 f.discharge_date, 
						 f.status, 
						 f.created_at
					 FROM 
					   financials f
					 INNER JOIN	
					   categories c on c.id = f.category_id
					 INNER JOIN	
						 persons p on p.id = f.person_id
					 WHERE 
				 	 	 LOWER(f.type) = $1 AND
						 LOWER(f.description) LIKE $2
	`
	financialList := []financials.Financial{}
	query := sqlg.NewQuery(context.Background())
	query.AddSQL(sql,
		strings.ToLower(filterType),
		strings.ToLower(description),
	)
	rows, err := query.Open()
	if err != nil {
		return &financialList, errors.New("Error listing Financials register: " + err.Error())
	}

	for rows.Next() {
		var financial financials.Financial
		if err := rows.Scan(
			&financial.ID,
			&financial.Type,
			&financial.Category.ID,
			&financial.Category.Type,
			&financial.Category.Description,
			&financial.Category.CreatedAt,
			&financial.Person.ID,
			&financial.Person.Name,
			&financial.Person.Customer,
			&financial.Person.Provider,
			&financial.Person.Document,
			&financial.Person.Cep,
			&financial.Person.Address,
			&financial.Person.State,
			&financial.Person.City,
			&financial.Person.Complement,
			&financial.Person.CreatedAt,
			&financial.Document,
			&financial.DocumentValue,
			&financial.DischargeValue,
			&financial.Description,
			&financial.DueDate,
			&financial.DischargeDate,
			&financial.Status,
			&financial.CreatedAt,
		); err != nil {
			return &financialList, errors.New("Error listing Descriptions: " + err.Error())
		}
		financialList = append(financialList, financial)
	}

	defer rows.Close()
	defer query.Close()
	return &financialList, nil
}

func (repository FinancialRepositoryDB) GetAllResume(filter financials.FinancialFilter) (*[]financials.FinancialResume, error) {
	filterType := filter.Type
	description := fmt.Sprintf("%%%s%%", filter.Description)

	sql := ` SELECT 
							f.id, 
							f.type, 
							f.category_id, 
							c.description,						 
							f.person_id, 
							p.name,						 						 
							f.document, 
							f.document_value,
							f.discharge_value,
							f.description, 
							f.due_date, 
							f.discharge_date, 
							f.status, 
							f.created_at
					 FROM 
					   financials f
					INNER JOIN 
						categories c ON c.id = f.category_id
					INNER JOIN 
						persons p ON p.id = f.person_id
				  WHERE 
					  LOWER(f.type) = $1 AND
					  LOWER(f.description) LIKE $2
	`
	financialList := []financials.FinancialResume{}
	query := sqlg.NewQuery(context.Background())
	query.AddSQL(sql,
		strings.ToLower(filterType),
		strings.ToLower(description),
	)
	rows, err := query.Open()
	if err != nil {
		return &financialList, errors.New("Error listing Financials register: " + err.Error())
	}

	for rows.Next() {
		var financial financials.FinancialResume
		if err := rows.Scan(
			&financial.ID,
			&financial.Type,
			&financial.Category.Id,
			&financial.Category.Description,
			&financial.Person.Id,
			&financial.Person.Name,
			&financial.Document,
			&financial.DocumentValue,
			&financial.DischargeValue,
			&financial.Description,
			&financial.DueDate,
			&financial.DischargeDate,
			&financial.Status,
			&financial.CreatedAt,
		); err != nil {
			return &financialList, errors.New("Error listing Descriptions: " + err.Error())
		}
		financialList = append(financialList, financial)
	}

	defer rows.Close()
	defer query.Close()
	return &financialList, nil
}

func (repository FinancialRepositoryDB) Update(financial *financials.Financial) error {
	query := `UPDATE financials
	          SET 
						  category_id = $1, 
							person_id = $2, 
							document = $3,			
							document_value = $4,
							description = $5,
							due_date = $6, 
							discharge_date = $7, 
							status = $8 
						WHERE 
						  ID = $9
	`
	if err := sqlg.Query(
		context.Background(), query,
		financial.Category.ID,
		financial.Person.ID,
		financial.Document,
		financial.DocumentValue,
		financial.Description,
		financial.DueDate,
		financial.DischargeDate,
		financial.Status,
		financial.ID).Execute(); err != nil {
		return errors.New("Error updating Financials register " + err.Error())
	}
	return nil
}

func (repository FinancialRepositoryDB) Cancel(financial *financials.Financial) error {
	query := `UPDATE financials
	          SET 
							updated_at = $1, 
							status = $2 
						WHERE 
						  ID = $3
	`
	if err := sqlg.Query(
		context.Background(), query,
		financial.UpdatedAt,
		financial.Status,
		financial.ID).Execute(); err != nil {
		return errors.New("Error Canceled Financials register " + err.Error())
	}
	return nil
}

func (repository FinancialRepositoryDB) Discharge(financial *financials.Financial) error {
	query := `UPDATE financials
	          SET
						  discharge_date = $1,  
						  discharge_value = $2,
							updated_at = $3, 
							status = $4 
						WHERE 
						  ID = $5
	`
	if err := sqlg.Query(
		context.Background(), query,
		financial.DischargeDate,
		financial.DischargeValue,
		financial.UpdatedAt,
		financial.Status,
		financial.ID).Execute(); err != nil {
		return errors.New("Error in discharge of Financial Registration " + err.Error())
	}
	return nil
}
