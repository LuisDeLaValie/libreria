package autormodels_test

import (
	"log"
	"os"
	"testing"
	"time"

	autormodels "github.com/TDTxLE/libreria/models/autor.models"
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

func TestCrearAutor(t *testing.T) {
	noww := time.Now()
	autor := autormodels.AutorModel{
		Nombre: "Autor de prueba",
		Creado: &noww,
	}

	nuevoAutor, err := autormodels.CrearAutor(autor)
	if err != nil {
		log.Fatalf("TestCrearAutor:%s", err.Error())
		t.Error(err.Error())
		t.Fail()
	}
	oid = nuevoAutor.Hex()
}

func TestListarAutores(t *testing.T) {

	_, err := autormodels.ListarAutores()
	if err != nil {
		log.Fatalf("TestListarAutores:%s", err.Error())
		t.Error(err.Error())
		t.Fail()
	}
}

func TestObtenerAutor(t *testing.T) {
	t.Run("libro existente", func(t *testing.T) {
		_, err := autormodels.ObtenerAutor("64aa30d37d3356b2bca23c65")
		if err != nil {
			log.Fatalf("TestObtenerAutor>libro existente:%s", err.Error())
			t.Fatal(err.Error())
			t.Fail()
		}

	})
	t.Run("libro no existe", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != nil {
				log.Fatalf("TestObtenerAutor>ibro no existe:%v", err)
				t.Error(err)
				t.Fail()
			}
		}()

		_, err := autormodels.ObtenerAutor("64a3b9bb60740a5a6707e647")
		if err == nil {
			log.Fatalf("TestObtenerAutor>ibro no existe:%s", "no genero error al escribir un id no exixtemte")
			t.Fatal("no genero error al escribir un id no exixtemte")
			t.FailNow()
		}
	})
	t.Run("oid mal escrito", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != nil {
				log.Fatalf("TestObtenerAutor>oid mal escrito:%v", err)
				t.Error(err)
				t.Fail()
			}
		}()

		_, err := autormodels.ObtenerAutor("64a3b9bb60740a5a6707")
		if err == nil {
			log.Fatalf("TestObtenerAutor>oid mal escrito:%s", "no genero error al escribir un id no valido")
			t.Fatal("no genero error al escribir un id no valido")
			t.FailNow()
		}

	})

}

func TestActualizarAutor(t *testing.T) {
	noww := time.Now()
	autor := autormodels.AutorModel{
		Nombre:      "Actualizar Autor",
		Actualizado: &noww,
	}

	err := autormodels.ActualizarAutor("64aa30d37d3356b2bca23c65", autor)
	if err != nil {
		log.Fatalf("TestActualizarAutor:%s", err.Error())
		t.Error(err.Error())
		t.Fail()
	}
}

func TestEliminarAutor(t *testing.T) {
	err := autormodels.EliminarAutor(oid)
	if err != nil {
		log.Fatalf("TestActualizarAutor:%s", err.Error())
		t.Errorf("Error al eliminar libro: %v", err)
		t.Fail()
	}
}
