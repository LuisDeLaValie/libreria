package autormodels_test

import (
	"testing"
	"time"

	autormodels "github.com/TDTxLE/libreria/models/autor.models"
)

var oid string

func TestCrearAutor(t *testing.T) {
	noww := time.Now()
	autor := autormodels.AutorModel{
		Nombre: "Autor de prueba",
		Creado: &noww,
	}

	nuevoAutor, err := autormodels.CrearAutor(autor)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
	oid = nuevoAutor.Hex()
}

func TestListarAutores(t *testing.T) {

	_, err := autormodels.ListarAutores()
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
}

func TestObtenerAutor(t *testing.T) {
	t.Run("libro existente", func(t *testing.T) {
		_, err := autormodels.ObtenerAutor("64aa30d37d3356b2bca23c65")
		if err != nil {
			t.Fatal(err.Error())
			t.Fail()
		}

	})
	t.Run("libro no existe", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != nil {
				t.Error(err)
				t.Fail()
			}
		}()

		_, err := autormodels.ObtenerAutor("64a3b9bb60740a5a6707e647")
		if err == nil {
			t.Fatal("no genero error al escribir un id no exixtemte")
			t.FailNow()
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

		_, err := autormodels.ObtenerAutor("64a3b9bb60740a5a6707")
		if err == nil {
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
		t.Error(err.Error())
		t.Fail()
	}
}

func TestEliminarAutor(t *testing.T) {
	err := autormodels.EliminarAutor(oid)
	if err != nil {
		t.Errorf("Error al eliminar libro: %v", err)
		t.Fail()
	}
}
