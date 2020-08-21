package main

import (
	"net/http"

	"github.com/GolangNorthwindRestApi/database"
	"github.com/GolangNorthwindRestApi/product"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	Databaseconnection := database.InitDB()
	defer Databaseconnection.Close()
	var productRepository = product.NewRepository(Databaseconnection)
	var productService product.Service
	productService = product.NewService(productRepository)
	r := chi.NewRouter()
	r.Mount("/products", product.MakeHttpHandler(productService))
	http.ListenAndServe(":3000", r)

}
