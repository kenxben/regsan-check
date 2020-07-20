package nso

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (h handler) buscarNSO(key string) (Producto, bool) {
	nso, ok := (*h.productos)[key]
	return nso, ok
}

func (h handler) buscarNombre(name string) (Producto, bool) {
	var p Producto
	return p, true
}

type handler struct {
	productos *TablaProductos
}

// NewHandler creates a handler with query functions to a product table.
func NewHandler(p *TablaProductos) http.Handler {
	return handler{p}
}

func writeResponse(w http.ResponseWriter, nso Producto) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	nsodata, err := json.Marshal(nso)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error al construir json"}`))
		log.Println("Error al construir json.", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"resultado": %s}`, nsodata)))
	return
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// vars := mux.Vars(r)
	key := r.FormValue("nso")
	name := r.FormValue("nombre")

	nso, ok := h.buscarNSO(key)
	if !ok {
		nso, ok = h.buscarNombre(name)
	}

	if ok {
		writeResponse(w, nso)
		return
	}

	// nso not found in data
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "NSO not found"}`))
	return
}
