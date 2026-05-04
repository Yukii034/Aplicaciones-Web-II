package cafeteria

import "testing"

func TestGuardar_DebeAlmacenarUnProducto(t *testing.T) {
	//Arrange - arreglar los datos
	repo := NewRepoMemoria()
	p := Producto{ID: 1, Nombre: "", Precio: 5.5, Stock: 10, Categoria: ""}

	//Act - ejecuta la funcion
	err := repo.GuardarProducto(p)

	//Assert - saber si se ha logrado correctamente
	if err != nil {
		t.Errorf("Error al guardar el producto %v", err)
	}

	if len(repo.ListarProductos()) != 1 {
		t.Errorf("El producto no se guardo en el repositorio")
	}
}

func TestBuscarPorID_DebeRetornarErrorSiNoExiste(t *testing.T) {
	//Arrange
	repo := NewRepoMemoria()
	//Act
	_, err := repo.ObtenerProducto(99)
	//Assert
	if err != nil {
		t.Errorf("Error al obtener el producto %v", err)
	}

	//Si se importa errors
	//if !errors.Is(err, err.ErrProductoNoEncontrado){ //Para saber si el error es igual a otro, no muestra algo redundante
	//	t.Errorf("Se esperaba el error ErrNoEncontrado, pero se obtuvo %v", err)
	//}
}
