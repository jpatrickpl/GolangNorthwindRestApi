package main

import (
	"net/http"

	"github.com/GolangNorthwindRestApi/customer"
	"github.com/GolangNorthwindRestApi/database"
	"github.com/GolangNorthwindRestApi/employee"
	"github.com/GolangNorthwindRestApi/product"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	Databaseconnection := database.InitDB()
	defer Databaseconnection.Close()
	var (
		productRepository  = product.NewRepository(Databaseconnection)
		employeeRepository = employee.NewRepository(Databaseconnection)
		customerRepository = customer.NewRepository(Databaseconnection)
	)
	var (
		productService  product.Service
		employeeService employee.Service
		customerService customer.Service
	)

	productService = product.NewService(productRepository)
	employeeService = employee.NewService(employeeRepository)
	customerService = customer.NewService(customerRepository)

	r := chi.NewRouter()
	r.Mount("/products", product.MakeHttpHandler(productService))
	r.Mount("/employees", employee.MakeHttpHandler(employeeService))
	r.Mount("/customers", customer.MakeHttpHandler(customerService))
	http.ListenAndServe(":3000", r)

}
