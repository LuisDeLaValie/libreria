package autor_test

import (
	"testing"
	"time"

	models "github.com/TDTxLE/libreria/models/autor"
)

var oid string

func TestCrearAutor(t *testing.T) {

	autor := models.AutorModel{
		Nombre: "Autor de prueba",
		Creado: time.Now(),
	}

	nuevoAutor, err := autor.CrearAutor(autor)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
	oid = nuevoAutor.Id.Hex()
}

func TestListarAutores(t *testing.T) {
	var lirbo models.AutorModel

	_, err := lirbo.ListarAutores()
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
}

func TestObtenerAutor(t *testing.T) {
	var libro models.AutorModel
	t.Run("libro existente", func(t *testing.T) {
		_, err := libro.ObtenerAutor("64aa30d37d3356b2bca23c65")
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

		_, err := libro.ObtenerAutor("64a3b9bb60740a5a6707e647")
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

		_, err := libro.ObtenerAutor("64a3b9bb60740a5a6707")
		if err == nil {
			t.Fatal("no genero error al escribir un id no valido")
			t.FailNow()
		}

	})

}

func TestActualizarAutor(t *testing.T) {
	autor := models.AutorModel{
		Nombre:      "Actualizar Autor",
		Actualizado: time.Now(),
	}

	err := autor.ActualizarAutor("64aa30d37d3356b2bca23c65", autor)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
}

func TestEliminarAutor(t *testing.T) {
	var autor models.AutorModel
	err := autor.EliminarAutor(oid)
	if err != nil {
		t.Errorf("Error al eliminar libro: %v", err)
		t.Fail()
	}
}
