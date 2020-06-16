package main


import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	// "verify-regsan-api/csvreader"
	"verify-regsan-api/gsheetsreader"
)


// var Productos map[string]csvreader.Producto
var Productos map[string]gsheetsreader.Producto


func init() {

	// Read data from spreadsheet using its id
	// Example sheet https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	Productos = gsheetsreader.GetData("1fM4UC4Y9uxkzJxIKFnW0h_YVAlB5qQ2cQdl4VuxmnnA")
}


func main() {

	// Serve api
	port := ":8080"
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/{NSO}/", buscarNSO).Methods(http.MethodGet)
	fmt.Printf("Serving at localhost%s", port)
	log.Fatal(http.ListenAndServe(port, r))
}