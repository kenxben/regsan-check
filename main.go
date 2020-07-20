package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regsan-check/nso"

	"github.com/gorilla/mux"
)

// map for product data using nso keys
var productos nso.TablaProductos

func init() {
	// initialize productos map
	// Read data from spreadsheet using its id
	sheetID := "1fM4UC4Y9uxkzJxIKFnW0h_YVAlB5qQ2cQdl4VuxmnnA"
	productos = nso.GetDataFromGsheet(sheetID)
}

func main() {

	// Serve api
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT not set, defaulting to 8080")
		port = "8080"
	}

	r := mux.NewRouter()
	h := nso.NewHandler(&productos)

	r.Queries(
		"nso", "[a-zA-z0-9-\\s]+",
		"nombre", "[a-zA-z0-9\\s]+",
	)
	r.Handle("/check-nso", h).Methods(http.MethodGet)
	fmt.Printf("Serving at port:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
