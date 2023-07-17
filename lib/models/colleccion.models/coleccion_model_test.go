package collecionmodels_test

import (
	"log"
	"os"
	"testing"
	"time"

	collecionmodels "github.com/TDTxLE/libreria/models/colleccion.models"
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

func TestCrearColeccion(t *testing.T) {

	titulo := "Coleccion de prueba"
	sipnosis := "Este es una coleccion de prueba para ver el funcionaminto dela api"
	creado := time.Now()
	coleccion := collecionmodels.ColleccionModel{Titulo: &titulo, Sinopsis: &sipnosis, Creado: creado}

	nuevoColeccion, err := collecionmodels.CrearColeccion(coleccion)
	if err != nil {
		log.Fatalf("TestCrearColeccion:\n\t%s", err.Error())
		t.Error(err.Error())
		t.Fail()
	}
	oid = nuevoColeccion.Hex()

}
func TestListarColecciones(t *testing.T) {

	_, err := collecionmodels.ListarColecciones()
	if err != nil {
		log.Fatalf("TestListarColecciones:\n\t%s", err.Error())
		t.Error(err.Error())
		t.Fail()
	}
}
func TestObtenerColeccion(t *testing.T) {
	t.Run("coleccion existente", func(t *testing.T) {
		_, err := collecionmodels.ObtenerColeccion("64aa468edce8053e1323c283")
		if err != nil {
			log.Fatalf("TestObtenerColeccion>coleccion existente:\n\t%s", err.Error())

			t.Fatal(err.Error())
			t.Fail()
		}

	})
	t.Run("coleccion no existe", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != nil {
				log.Fatalf("TestObtenerColeccion>coleccion no existente:\n\t%v", err)
				t.Error(err)
				t.Fail()
			}
		}()

		_, err := collecionmodels.ObtenerColeccion("64a3b9bb60740a5a6707e647")
		if err == nil {
			log.Fatalf("TestObtenerColeccion>coleccion no existente:\n\t%s", "no se creeo el error")
			t.Error("no se creeo el error ")
			t.Fail()
		}
	})
	t.Run("oid mal escrito", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != nil {
				log.Fatalf("TestObtenerColeccion>oid mal escrito:\n\t%v", err)
				t.Error(err)
				t.Fail()
			}
		}()

		_, err := collecionmodels.ObtenerColeccion("64a3b9bb60740a5a6707")
		if err == nil {
			log.Fatalf("TestObtenerColeccion>oid mal escrito:\n\t%s", "no se creeo el error ")
			t.Error("no se creeo el error ")
			t.Fail()
		}

	})
}
func TestActualizarColeccion(t *testing.T) {
	clock := time.Now()
	auxTitulo := "actualizar titulo " + clock.Local().GoString()

	var libro collecionmodels.ColleccionModel
	libro.Titulo = &auxTitulo
	libro.Actualizado = clock

	err := collecionmodels.ActualizarColeccion("64aa468edce8053e1323c283", libro)
	if err != nil {
		log.Fatalf("TestActualizarColeccion:\n\t%s", err.Error())
		t.Error(err.Error())
		t.Fail()
	}
}
func TestEliminarColeccion(t *testing.T) {
	err := collecionmodels.EliminarColeccion(oid)
	if err != nil {
		log.Fatalf("TestEliminarColeccion:\n\t%s", err.Error())
		t.Error(err.Error())
		t.Fail()
	}
}
