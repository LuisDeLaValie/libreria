package librosmodels_test

import (
	"log"
	"os"
	"testing"
	"time"

	librosmodels "github.com/TDTxLE/libreria/models/libros.models"
	utils "github.com/TDTxLE/libreria/utils/mongoTogolang"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMain(m *testing.M) {
	// Abrir el archivo de registro en modo de añadir (append)
	var err error
	logFile, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Configurar el logger para que escriba en el archivo de registro
	log.SetOutput(logFile)

	// Ejecutar los tests
	code := m.Run()

	// Cerrar el archivo de registro
	logFile.Close()

	// Finalizar los tests
	os.Exit(code)
}

var oid string

func TestCrearLibro(t *testing.T) {

	titulo := "Lirbo de prueba"
	sipnosis := "Este es un libro de prueba para ver el funcionaminto dela api"
	creado := time.Now()
	libro := librosmodels.LibroModelForm{Titulo: &titulo, Sinopsis: &sipnosis, Creado: creado}

	nuevoLibro, err := librosmodels.CrearLibro(libro)
	if err != nil {
		log.Fatalf("TestCrearLibro: No se pudo crear el libro: %v", err)
		t.Errorf("No se pudo crear el libro: %v", err)
		t.Fail()
	}
	oid = nuevoLibro.Hex()
	log.Printf("TestCrearLibro: Libro creado con exito : %s", nuevoLibro)
}

func TestListarLibros(t *testing.T) {

	_, err := librosmodels.ListarLibros()

	if err != nil {
		log.Fatalf("TestListarLibros: Error al obtener los libros: %v", err)
		t.Errorf("Error al obtener los libros: %v", err)
		t.Fail()
	}
}

func TestObtenerLibro(t *testing.T) {

	t.Run("libro existente", func(t *testing.T) {

		libroobtenido, err := librosmodels.ObtenerLibro("64a3b9bb60740a5a6707e64b")
		if err != nil {
			log.Fatalf("TestObtenerLibro > libro existente: Error al obtenr el libro: %v", err)
			t.Errorf("Error al obtenr el libro: %v", err)
			t.Fail()
		}
		log.Printf("TestObtenerLibro > libro existente: libro obtenido: %v", libroobtenido)

	})
	t.Run("libro no existe", func(t *testing.T) {

		defer func() {
			err := recover()
			if err != nil {
				log.Fatalf("TestObtenerLibro > libro no existe: %v", err)
				t.Error(err)
				t.Fail()
			}
		}()

		_, err := librosmodels.ObtenerLibro("64a3b9bb60740a5a6707e647")
		if err != nil {
			log.Printf("TestObtenerLibro > libro no existe: Error al obtenr el libro: %v", err)
			t.Logf("Error al obtenr el libro: %v", err)
		}
	})
	t.Run("oid mal escrito", func(t *testing.T) {

		defer func() {
			err := recover()
			if err != nil {
				log.Fatal(err)
				t.Error(err)
				t.Fail()
			}
		}()

		_, err := librosmodels.ObtenerLibro("64a3b9bb60740a5a6707")
		if err == nil {
			log.Fatalf("TestObtenerLibro > oid mal escrito: Error al obtenr el libro: %v", err)
		}

	})
}

func TestActualizarLibro(t *testing.T) {

	oidc, _ := utils.ValidarOID("64aa468edce8053e1323c283")
	coleccion := oidc
	oid, _ := utils.ValidarOID("64aa30d37d3356b2bca23c65")
	autores := []primitive.ObjectID{*oid}
	clock := time.Now()
	auxTitulo := "actualizar titulo " + clock.Local().GoString()

	var libro librosmodels.LibroModelForm
	libro.Titulo = &auxTitulo
	libro.Actualizado = clock
	libro.Autores = &autores
	libro.Collection = *coleccion

	err := librosmodels.ActualizarLibro("64a5cf610d0195d943d2ee10", libro)
	if err != nil {
		log.Fatalf("TestActualizarLibro: Error al actualizar libro: %v", err)
		t.Errorf("Error al actualizar libro: %v", err)
		t.Fail()

	}
}

func TestEliminarLibro(t *testing.T) {

	err := librosmodels.EliminarLibro(oid)
	if err != nil {
		log.Fatalf("TestEliminarLibro: Error al eliminar libro: %v", err)
		t.Errorf("Error al eliminar libro: %v", err)
		t.Fail()

	}
}

func TestActualizarVariosLibros(t *testing.T) {

	id, _ := utils.ValidarOID("64ab220b086c781da8d8b1e9")
	oids := []string{
		"64ab21286849d0fc826cc1f1",
		"64ab220b086c781da8d8b1e9",
	}
	// titulo := "Actualizacion multople"
	var actu librosmodels.LibroModelForm
	// actu.Titulo = &titulo
	actu.Collection = *id
	err := librosmodels.ActualizarVariosLibros(oids, actu)

	if err != nil {
		log.Fatalf("TestEliminarLibro: %v", err)
		t.Error(err)
		t.Fail()
	}
}
