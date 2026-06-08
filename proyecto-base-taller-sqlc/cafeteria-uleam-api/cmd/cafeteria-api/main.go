// Command cafeteria-api arranca el servidor HTTP de la Cafetería Universitaria.
package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/glebarez/sqlite" // driver SQLite pure-Go (sin CGO)
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"cafeteria-uleam-api/internal/handlers"
	"cafeteria-uleam-api/internal/middleware"
	"cafeteria-uleam-api/internal/models"
	"cafeteria-uleam-api/internal/storage"
)

func main() {
	// 1. Abrir SQLite y migrar el esquema (crea las tablas si no existen).
	db, err := gorm.Open(sqlite.Open("cafeteria.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}
	if err := db.AutoMigrate(&models.Producto{}, &models.Categoria{}); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}

	// 2. Crear el almacenamiento GORM y sembrar si está vacío.
	almacenGorm := storage.NuevoAlmacenSQLite(db)
	almacenGorm.SembrarSiVacio()

	var almacen storage.Almacen
	switch os.Getenv("STORAGE") {
	case "sqlc":
		if sqlDB, err := db.DB(); err == nil {
			_ = sqlDB.Close() // GORM ya migró y sembró; liberamos su conexión
		}
		sdb, err := sql.Open("sqlite", "cafeteria.db")
		if err != nil {
			log.Fatal("sqlc: ", err)
		}
		almacen = storage.NuevoAlmacenSQLC(sdb)
	default:
		almacen = almacenGorm
	}

	// 3. Server con inyección de dependencias (recibe un storage.Almacen).
	servidor := handlers.NewServer(almacen) // idéntico: no sabe cuál recibió

	// 4. Router + middleware.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 5. Rutas versionadas /api/v1/.
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/productos", servidor.ListarProductos)
		r.Post("/productos", servidor.CrearProducto)
		r.Get("/productos/{id}", servidor.ObtenerProducto)
		r.Put("/productos/{id}", servidor.ActualizarProducto)
		r.Delete("/productos/{id}", servidor.BorrarProducto)

		r.Get("/categorias", servidor.ListarCategorias)
		r.Post("/categorias", servidor.CrearCategoria)
		r.Get("/categorias/{id}", servidor.ObtenerCategoria)
		r.Put("/categorias/{id}", servidor.ActualizarCategoria)
		r.Delete("/categorias/{id}", servidor.BorrarCategoria)
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
