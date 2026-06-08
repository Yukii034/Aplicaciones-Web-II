// Command cafeteria-api arranca el servidor HTTP de la Cafetería Universitaria.
package main

import (
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"cafeteria-uleam-api/internal/handlers"
	"cafeteria-uleam-api/internal/middleware"
	"cafeteria-uleam-api/internal/models"
	"cafeteria-uleam-api/internal/storage"
)

func main() {
	db, err := gorm.Open(sqlite.Open("cafeteria.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos", err)
	}
	if err := db.AutoMigrate(&models.Producto{}, &models.Categoria{}); err != nil {
		log.Fatal("Error al migrar:", err)
	}

	// 1. Crear el almacenamiento y cargar datos iniciales.
	almacen := storage.NewAlmacenSQLite(db)
	almacen.SembrarVacio()

	// 2. Crear el Server inyectándole el almacenamiento.
	servidor := handlers.NewServer(almacen)

	// 3. Configurar el router.
	r := chi.NewRouter()

	// 4. Middleware GLOBAL: se aplica a TODAS las peticiones, en orden.
	//    Logger    -> registra método, ruta y tiempo de cada request.
	//    Recoverer -> si un handler entra en panic, responde 500 en vez de tumbar el server.
	//    CORS      -> nuestro middleware propio (ver internal/middleware).
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 5. Rutas versionadas /api/v1/.
	r.Route("/api/v1", func(r chi.Router) {
		// Productos: CRUD completo.
		r.Get("/productos", servidor.ListarProductos)
		r.Post("/productos", servidor.CrearProducto)
		r.Get("/productos/{id}", servidor.ObtenerProducto)
		r.Put("/productos/{id}", servidor.ActualizarProducto)
		r.Delete("/productos/{id}", servidor.BorrarProducto)

		// Categorías: CRUD completo.
		r.Get("/categorias", servidor.ListarCategorias)
		r.Post("/categorias", servidor.CrearCategoria)
		r.Get("/categorias/{id}", servidor.ObtenerCategoria)
		r.Put("/categorias/{id}", servidor.ActualizarCategoria)
		r.Delete("/categorias/{id}", servidor.BorrarCategoria)
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
