package main


import (
    "fmt"
    // "log"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)


func buscarNSO(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    w.Header().Set("Content-Type", "application/json; charset=utf-8")

    // search NSO in query
    if key, ok := vars["NSO"]; ok {

        // search NSO id (nso) in data
        if nso, ok := Productos[key]; ok {

        	nsodata, err := json.Marshal(nso)
	        if err != nil {
	            w.WriteHeader(http.StatusInternalServerError)
	            w.Write([]byte(`{"error": "error al construir json"}`))
	            return
	        }
        	w.WriteHeader(http.StatusOK)
	        w.Write([]byte(fmt.Sprintf(`{"resultado": %s}`, nsodata)))
	        return

	    // nso not found in data
        } else {
        	w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "Registro no encontrado en base"}`))
			return
        }
    }

    w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "NSO not found"}`))
}