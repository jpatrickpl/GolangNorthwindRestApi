package main

import (
	"database/sql"
	"encoding/json"
	"golangnorthwindrestapi/database"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

var Databaseconnection *sql.DB

type Product struct {
	ID           int    `json:"id"`
	Product_Code string `json:"product_code"`
	Description  string `json:"description"`
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	Databaseconnection = database.InitDB()
	defer Databaseconnection.Close()

	r := chi.NewRouter()

	r.Get("/products", Allproductos)
	r.Post("/products", CreateProducto)
	r.Put("/products/{id}", UpdateProducto)
	r.Delete("/products/{id}", DeleteProducto)
	http.ListenAndServe(":3000", r)

}

func DeleteProducto(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	query, err := Databaseconnection.Prepare("delete from Products  where id = ?")
	catch(err)
	_, er := query.Exec(id)
	catch(er)
	defer query.Close()
	responsewithJSON(w, http.StatusCreated, map[string]string{"message": "succesfully deleted"})
}

func CreateProducto(w http.ResponseWriter, r *http.Request) {
	var producto Product
	json.NewDecoder(r.Body).Decode(&producto)
	query, err := Databaseconnection.Prepare("Insert Products set product_code =? , description =?")
	catch(err)
	_, er := query.Exec(producto.Product_Code, producto.Description)
	catch(er)
	defer query.Close()
	responsewithJSON(w, http.StatusCreated, map[string]string{"message": "succesfully created"})
}

func UpdateProducto(w http.ResponseWriter, r *http.Request) {
	var producto Product
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&producto)

	query, err := Databaseconnection.Prepare("update Products set product_code =? , description =? where id = ?")
	catch(err)
	_, er := query.Exec(producto.Product_Code, producto.Description, id)
	catch(er)
	defer query.Close()
	responsewithJSON(w, http.StatusCreated, map[string]string{"message": "succesfully updated"})
}

func Allproductos(w http.ResponseWriter, r *http.Request) {

	const sql = "select id,product_code,coalesce(description,'') from products"
	results, err := Databaseconnection.Query(sql)
	catch(err)
	var products []*Product

	for results.Next() {
		product := &Product{}
		err := results.Scan(&product.ID, &product.Product_Code, &product.Description)
		catch(err)
		products = append(products, product)
	}
	responsewithJSON(w, http.StatusOK, products)
}

func responsewithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-type", "apllication/json")
	w.WriteHeader(code)
	w.Write(response)

}
