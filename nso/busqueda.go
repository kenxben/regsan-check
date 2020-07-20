package nso

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sahilm/fuzzy"
)

func (h handler) buscarCodigo(key string) ([]Producto, bool) {
	nso, ok := (*h.productos)[key]
	return []Producto{nso}, ok
}

func (h handler) buscarNombre(name string) ([]Producto, bool) {
	// slice to fill results
	var res []Producto

	// initialize and fill columns with keys and product names
	var desc []string // product description string based on name, brand, etc.
	var keys []string
	for k, v := range *h.productos {
		keys = append(keys, k)
		desc = append(
			desc,
			v.NombreProducto+" "+v.MarcaProducto+" "+v.Titular,
		)
	}

	// get matches with fuzzy search
	matches := fuzzy.Find(name, desc)

	// truncate to the first results
	if len(matches) > 5 {
		matches = matches[:5]
	}

	// build list of results
	for _, m := range matches {
		if m.Score > 0 {
			res = append(res, (*h.productos)[keys[m.Index]])
		}
	}

	// return ok if list is populated
	if len(res) > 0 {
		return res, true
	}

	// else return not ok
	return res, false
}

type handler struct {
	productos *TablaProductos
}

// NewHandler creates a handler with query functions to a product table.
func NewHandler(p *TablaProductos) http.Handler {
	return handler{p}
}

func writeResponse(w http.ResponseWriter, nso []Producto) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	nsodata, err := json.Marshal(nso)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error al construir json"}`))
		log.Println("Error al construir json.", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"resultados": %s, "metadata":{}}`, nsodata)))
	return
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// vars := mux.Vars(r)
	key := r.FormValue("nso")
	name := r.FormValue("producto")

	// try matching nso code
	nso, ok := h.buscarCodigo(key)
	if !ok {
		// if code didn't work, try by fuzzy search
		nso, ok = h.buscarNombre(name)
	}

	// if any search went ok, write results to response
	if ok {
		writeResponse(w, nso)
		return
	}

	// if not, search wasn't found in data
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "NSO no encontrado"}`))
	return
}
