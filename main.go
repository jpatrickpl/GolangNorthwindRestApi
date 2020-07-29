package main

import (
	"GolangNorthwindRestApi/database"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/*
		r := chi.NewRouter()
		r.Use(middleware.Logger)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome"))
		})
		http.ListenAndServe(":3000", r)*/

	Databaseconnection := database.InitDB()
	defer Databaseconnection.Close()
	fmt.Println(Databaseconnection)

}
