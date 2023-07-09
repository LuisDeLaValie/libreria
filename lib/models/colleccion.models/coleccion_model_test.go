package collecionmodels_test

import (
	"testing"
	"time"

	collecionmodels "github.com/TDTxLE/libreria/models/colleccion.models"
)

var oid string

func TestCrearColeccion(t *testing.T) {

	titulo := "Coleccion de prueba"
	sipnosis := "Este es una coleccion de prueba para ver el funcionaminto dela api"
	creado := time.Now()
	coleccion := collecionmodels.ColleccionModel{Titulo: &titulo, Sinopsis: &sipnosis, Creado: creado}

	nuevoColeccion, err := collecionmodels.CrearColeccion(coleccion)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
	oid = nuevoColeccion.Hex()

}

func TestListarColecciones(t *testing.T) {

	_, err := collecionmodels.ListarColecciones()
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
}

func TestObtenerColeccion(t *testing.T) {
	t.Run("coleccion existente", func(t *testing.T) {
		_, err := collecionmodels.ObtenerColeccion("64aa468edce8053e1323c283")
		if err != nil {
			t.Fatal(err.Error())
			t.Fail()
		}

	})
	t.Run("coleccion no existe", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != nil {
				t.Error(err)
				t.Fail()
			}
		}()

		_, err := collecionmodels.ObtenerColeccion("64a3b9bb60740a5a6707e647")
		if err == nil {
			t.Error("no se creeo el error ")
			t.Fail()
		}
	})
	t.Run("oid mal escrito", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != nil {
				t.Error(err)
				t.Fail()
			}
		}()

		_, err := collecionmodels.ObtenerColeccion("64a3b9bb60740a5a6707")
		if err == nil {
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
		t.Error(err.Error())
		t.Fail()
	}
}

func TestEliminarColeccion(t *testing.T) {
	err := collecionmodels.EliminarColeccion(oid)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
}
