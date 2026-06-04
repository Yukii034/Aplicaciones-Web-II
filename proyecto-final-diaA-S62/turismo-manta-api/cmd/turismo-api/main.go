// Command turismo-api arranca el servidor HTTP de la API de turismo Manta.
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"turismo-manta-api/internal/handlers"
	"turismo-manta-api/internal/middleware"
	"turismo-manta-api/internal/storage"
)

func main() {
	// 1. Crear el almacenamiento y cargar datos iniciales.
	almacen := storage.NuevaMemoria()
	almacen.Seed()

	// 2. Crear el Server inyectándole el almacenamiento.
	//    Aquí ocurre la inyección de dependencias: el Server NO crea
	//    su propio storage, lo RECIBE desde afuera.
	servidor := handlers.NewServer(almacen)

	// 3. Configurar el router con versionado /api/v1/.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	r.Route("/api/v1", func(r chi.Router) {
		// Negocios: handlers son métodos de Server (DI).
		r.Get("/negocios", servidor.ListarNegocios)
		r.Post("/negocios", servidor.CrearNegocio)
		r.Get("/negocios/{id}", servidor.ObtenerNegocio)
		r.Put("/negocios/{id}", servidor.ActualizarNegocio)
		r.Delete("/negocios/{id}", servidor.BorrarNegocio)

		// Turistas: handler suelto, asimétrico a propósito (S7 lo refactoriza).
		r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
			panic("panic de prueba")
		})

		r.Get("/turistas", handlers.ListarTuristas)
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
