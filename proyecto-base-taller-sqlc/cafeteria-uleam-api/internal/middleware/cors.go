package middleware

import "net/http"

// CORS habilita el consumo de la API desde cualquier origen (configuración
// permisiva, pensada para desarrollo y para las demos en clase).
//
// Fíjense en la forma: recibe `next` (el siguiente eslabón de la cadena) y
// devuelve un http.HandlerFunc que primero pone los headers de CORS y luego
// llama a `next`. Ese es EL patrón de todo middleware.
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// El navegador manda una petición OPTIONS de "preflight" antes de la real.
		// La respondemos de inmediato, sin pasar al handler.
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
