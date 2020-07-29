package database
import "database/sql"


func InitDB() *sql.DB {

	Connectionstring := "root:admin@tcp(localhost:3307)/northwind"
	Databaseconnection, err := sql.Open("mysql", Connectionstring)
	
	if err != nil {
		panic(err.Error())
	}
	return Databaseconnection
}
