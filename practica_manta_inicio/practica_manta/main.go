// =============================================================================
// PRÁCTICA GUIADA · SEMANA 3 · DÍA A
// =============================================================================
// Proyecto: Manta Tourist Connect
// Estado: punto de partida (código monolítico, sin paquetes, sin errores
//         idiomáticos, sin interfaces)
//
// TU TAREA (60 min): refactorizar este código aplicando los 3 conceptos de hoy.
// Revisa el handout para ver los 5 hitos que debes completar.
//
// Para ejecutar:   go run .
// =============================================================================

package main

import "fmt"

// -----------------------------------------------------------------------------
// LAS 5 ENTIDADES DE MANTA TOURIST CONNECT
// -----------------------------------------------------------------------------

type Crucero struct {
	ID        int
	Nombre    string
	Bandera   string
	Pasajeros int
}

type Turista struct {
	ID        int
	Nombre    string
	Idioma    string // "es", "en", "zh", "fr", ...
	CruceroID int    // referencia por ID
}

type Negocio struct {
	ID          int
	Nombre      string
	Tipo        string // "restaurante", "tour", "artesanía", ...
	IdiomasHabl []string
}

type Empleado struct {
	ID        int
	Nombre    string
	NegocioID int // referencia por ID
	Idiomas   []string
}

type CheckIn struct {
	ID         int
	TuristaID  int
	NegocioID  int
	EmpleadoID int
	CodigoQR   string
}

// -----------------------------------------------------------------------------
// "BASE DE DATOS" EN MEMORIA
// -----------------------------------------------------------------------------

var cruceros = []Crucero{}
var turistas = []Turista{}
var negocios = []Negocio{}
var empleados = []Empleado{}
var checkins = []CheckIn{}

var siguienteCheckInID = 1

// -----------------------------------------------------------------------------
// FUNCIONES — todas tienen el mismo problema: fracasan silenciosamente
// retornando zero values cuando algo no existe.
// -----------------------------------------------------------------------------

func AgregarCrucero(c Crucero)   { cruceros = append(cruceros, c) }
func AgregarTurista(t Turista)   { turistas = append(turistas, t) }
func AgregarNegocio(n Negocio)   { negocios = append(negocios, n) }
func AgregarEmpleado(e Empleado) { empleados = append(empleados, e) }

func BuscarCruceroPorID(id int) Crucero {
	for _, c := range cruceros {
		if c.ID == id {
			return c
		}
	}
	return Crucero{} // ← FRACASO SILENCIOSO (arreglar en Hito 2)
}

func BuscarTuristaPorID(id int) Turista {
	for _, t := range turistas {
		if t.ID == id {
			return t
		}
	}
	return Turista{}
}

func BuscarNegocioPorID(id int) Negocio {
	for _, n := range negocios {
		if n.ID == id {
			return n
		}
	}
	return Negocio{}
}

func BuscarEmpleadoPorID(id int) Empleado {
	for _, e := range empleados {
		if e.ID == id {
			return e
		}
	}
	return Empleado{}
}

// RegistrarCheckIn es la función CLAVE para el ejercicio.
// Puede fallar por múltiples razones:
//   - el turista no existe
//   - el negocio no existe
//   - el empleado no existe o no trabaja en ese negocio
//   - el empleado no habla el idioma del turista
//
// Actualmente: si algo falla, simplemente no registra nada y no avisa cuál
// fue la causa. En Hito 3 tendrás que reportar CUÁL razón causó el problema.
func RegistrarCheckIn(turistaID, negocioID, empleadoID int, qr string) CheckIn {
	turista := BuscarTuristaPorID(turistaID)
	if turista.ID == 0 {
		fmt.Println("algo salió mal...") // ← inútil, no dice qué
		return CheckIn{}
	}

	negocio := BuscarNegocioPorID(negocioID)
	if negocio.ID == 0 {
		fmt.Println("algo salió mal...")
		return CheckIn{}
	}

	empleado := BuscarEmpleadoPorID(empleadoID)
	if empleado.ID == 0 || empleado.NegocioID != negocio.ID {
		fmt.Println("algo salió mal...")
		return CheckIn{}
	}

	// ¿El empleado habla el idioma del turista?
	hablaIdioma := false
	for _, i := range empleado.Idiomas {
		if i == turista.Idioma {
			hablaIdioma = true
			break
		}
	}
	if !hablaIdioma {
		fmt.Println("algo salió mal...")
		return CheckIn{}
	}

	// Todo bien: registramos
	c := CheckIn{
		ID:         siguienteCheckInID,
		TuristaID:  turistaID,
		NegocioID:  negocioID,
		EmpleadoID: empleadoID,
		CodigoQR:   qr,
	}
	siguienteCheckInID++
	checkins = append(checkins, c)
	return c
}

// -----------------------------------------------------------------------------
// MAIN — datos de prueba y llamadas
// -----------------------------------------------------------------------------

func main() {
	// Cruceros en el puerto de Manta
	AgregarCrucero(Crucero{ID: 1, Nombre: "MS Zaandam", Bandera: "Países Bajos", Pasajeros: 1432})

	// Turistas a bordo
	AgregarTurista(Turista{ID: 10, Nombre: "Emma Schneider", Idioma: "en", CruceroID: 1})
	AgregarTurista(Turista{ID: 11, Nombre: "Li Wei", Idioma: "zh", CruceroID: 1})

	// Negocios locales
	AgregarNegocio(Negocio{ID: 100, Nombre: "Marisquería La Playa", Tipo: "restaurante", IdiomasHabl: []string{"es", "en"}})
	AgregarNegocio(Negocio{ID: 101, Nombre: "Tours Manta Aventura", Tipo: "tour", IdiomasHabl: []string{"es"}})

	// Empleados
	AgregarEmpleado(Empleado{ID: 1000, Nombre: "María", NegocioID: 100, Idiomas: []string{"es", "en"}})
	AgregarEmpleado(Empleado{ID: 1001, Nombre: "Pedro", NegocioID: 101, Idiomas: []string{"es"}})

	// CASO 1: Check-in exitoso (María habla inglés, Emma habla inglés)
	c := RegistrarCheckIn(10, 100, 1000, "QR-001")
	fmt.Printf("Check-in 1: %+v\n\n", c)

	// CASO 2: Pedro no habla chino → debería reportar ErrIdiomaNoCoincide
	c = RegistrarCheckIn(11, 101, 1001, "QR-002")
	fmt.Printf("Check-in 2: %+v\n\n", c)

	// CASO 3: turista 999 no existe → debería reportar ErrTuristaNoEncontrado
	c = RegistrarCheckIn(999, 100, 1000, "QR-003")
	fmt.Printf("Check-in 3: %+v\n\n", c)

	// CASO 4: empleado 1000 no trabaja en el negocio 101
	c = RegistrarCheckIn(10, 101, 1000, "QR-004")
	fmt.Printf("Check-in 4: %+v\n\n", c)
}
