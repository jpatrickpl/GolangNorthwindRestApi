package employee

import (
	"database/sql"

	"github.com/GolangNorthwindRestApi/helper"
)

type Repository interface {
	GetEmployees(params *getEmployeesRequest) ([]*Employee, error)
	GetTotalEmployees() (int64, error)
	GetEmployeeById(param *getEmployeeByIdRequest) (*Employee, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (repo *repository) GetEmployees(params *getEmployeesRequest) ([]*Employee, error) {
	const sql = `SELECT id, first_name,last_name,
				 company, email_address, job_title,
				 business_phone,home_phone,
				 coalesce (mobile_phone,''), fax_number,address
				 FROM EMPLOYEES
				 LIMIT ? OFFSET ?`

	results, err := repo.db.Query(sql, params.Limit, params.Offset)
	helper.Catch(err)

	var employees []*Employee

	for results.Next() {
		employee := &Employee{}

		err = results.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Company,
			&employee.EmailAddress, &employee.JobTitle, &employee.BusinessPhone, &employee.HomePhone,
			&employee.MobilePhone, &employee.FaxNumber, &employee.Address)
		helper.Catch((err))
		employees = append(employees, employee)

	}

	return employees, nil

}

func (repo *repository) GetTotalEmployees() (int64, error) {
	const sql = `SELECT COUNT(*) FROM EMPLOYEES`
	var total int64

	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch((err))
	return total, err

}

func (repo *repository) GetEmployeeById(param *getEmployeeByIdRequest) (*Employee, error) {
	const sql = `SELECT id, first_name,last_name,
	company, email_address, job_title,
	business_phone,home_phone,
	coalesce (mobile_phone,''), fax_number,address
	FROM EMPLOYEES
	where id = ?`
	row := repo.db.QueryRow(sql, param.EmployeeID)
	employee := &Employee{}

	err := row.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Company,
		&employee.EmailAddress, &employee.JobTitle, &employee.BusinessPhone, &employee.HomePhone,
		&employee.MobilePhone, &employee.FaxNumber, &employee.Address)
	helper.Catch((err))
	return employee, nil
}

func (s *service) GetEmployeeById(param *getEmployeeByIdRequest) (*Employee, error) {
	return s.repo.GetEmployeeById(param)
}
