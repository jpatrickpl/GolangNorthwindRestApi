package employee

import (
	"database/sql"

	"github.com/GolangNorthwindRestApi/helper"
)

type Repository interface {
	GetEmployees(params *getEmployeesRequest) ([]*Employee, error)
	GetTotalEmployees() (int64, error)
	GetEmployeeById(param *getEmployeeByIdRequest) (*Employee, error)
	GetBestEmployee() (*BestEmployee, error)
	InsertEmployee(params *addEmployeeRequest) (int64,error)
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

func (repo *repository) GetBestEmployee() (*BestEmployee, error) {
	const sql = `SELECT e.id, count(e.id) as totalVentas,e.first_name,e.last_name
	FROM orders o
	inner join employees e on o.employee_id = e.id
	group by o.employee_id
	order by totalVentas desc
	limit 1`

	row := repo.db.QueryRow(sql)
	employee := &BestEmployee{}
	err := row.Scan(&employee.ID, &employee.TotalVentas, &employee.LastName, &employee.FirstName)
	helper.Catch(err)

	return employee, nil

}

func (repo *repository) InsertEmployee(params *addEmployeeRequest) (int64,error) {
	const sql = `insert into EMPLOYEES   ( first_name,last_name,
	company, address,business_phone,email_address,fax_number,home_phone, job_title,mobile_phone ) values(?,?,?,?,?,?,?,?,?,?)`

	result, err := repo.db.Exec(sql,params.FirstName,params.LastName,params.Company,params.Address,
			 params.BusinessPhone,params.EmailAddress,params.FaxNumber,params.HomePhone,params.JobTitle,params.MobilePhone)

			 helper.Catch(err)
			 id,err:=result.LastInsertId()
			 helper.Catch(err)
			 return id,nil
}