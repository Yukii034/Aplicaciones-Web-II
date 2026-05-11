package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Negocio struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Tipo   string `json:"tipo"`
	Ciudad string `json:"ciudad"`
}

var (
	mux      sync.Mutex
	negocios = []Negocio{
		{ID: 1, Nombre: "Hotel Online", Tipo: "Hotel", Ciudad: "Madrid"},
		{ID: 2, Nombre: "Hotel Online", Tipo: "Hotel", Ciudad: "Madrid"},
		{ID: 3, Nombre: "Hotel Online", Tipo: "Hotel", Ciudad: "Madrid"},
	}
	siguienteID int = 4
)

func main() {
	http.HandleFunc("negocios/", manejarNegocio)
	http.ListenAndServe(":8080", nil)
	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola Mundo") //Response (w ResponseWriter)
	})
	http.HandleFunc("/saludo", func(w http.ResponseWriter, r *http.Request) {
		nombre := r.URL.Query().Get("nombre")
		if nombre == "" {
			nombre = "Pierina"
		}
		fmt.Fprintf(w, "Hola %s!", nombre)
	})*/
	/*http.HandleFunc("/negocios", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", "GET")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		negocios := []Negocio{
			Negocio{ID: 1, Nombre: "Hotel Online", Tipo: "Hotel", Ciudad: "Madrid"},
			Negocio{ID: 2, Nombre: "Hotel Online", Tipo: "Hotel", Ciudad: "Madrid"},
			Negocio{ID: 3, Nombre: "Hotel Online", Tipo: "Hotel", Ciudad: "Madrid"},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(negocios) //http://localhost:8080/negocios
	})
	http.ListenAndServe(":8080", nil)*/
}

func crearNegocio(w http.ResponseWriter, r *http.Request) {
	var nuevo Negocio
	if error := json.NewDecoder(r.body).Decode(&nuevo); error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}
	mux.Lock()
	nuevo.ID = siguienteID
	siguienteID++
	negocios = append(negocios, nuevo)
	mux.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevo)
}

func listarNegocio(w http.ResponseWriter, r *http.Request) {
	mux.Lock()
	defer mux.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(negocios)
}

func manejarNegocios(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listarNegocio(w, r)
	case http.MethodPost:
		crearNegocio(w, r)
	default:
		w.Header().Set("Allow", "GET", "POST")
	}
}
