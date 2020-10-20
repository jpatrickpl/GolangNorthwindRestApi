package customer

import (
	"database/sql"

	"github.com/GolangNorthwindRestApi/helper"
)

type Repository interface {
	GetCustomers(param *getCustomerRequest) ([]*Customer, error)
	GetTotalCustomers() (int64, error)
}
type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (repo repository) GetTotalCustomers() (int64, error) {
	const sql = "select count(*) from customers"
	var total int64
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)
	return total, nil
}

func (repo repository) GetCustomers(param *getCustomerRequest) ([]*Customer, error) {

	const sql = "select id, first_Name,   last_Name,      address,  business_phone    ,city,  company   from customers c limit ? offset ?"

	results, err := repo.db.Query(sql, param.Limit, param.Offset)
	helper.Catch(err)
	var customers []*Customer

	for results.Next() {
		customer := &Customer{}

		err = results.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Address, &customer.BussinessPhone, &customer.City, &customer.Company)
		helper.Catch(err)
		customers = append(customers, customer)
	}

	return customers, nil

}
