package librosmodels_test

import (
	"testing"
	"time"

	librosmodels "github.com/TDTxLE/libreria/models/libros.models"
	utils "github.com/TDTxLE/libreria/utils/mongoTogolang"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var oid string

func TestCrearLibro(t *testing.T) {

	titulo := "Lirbo de prueba"
	sipnosis := "Este es un libro de prueba para ver el funcionaminto dela api"
	creado := time.Now()
	libro := librosmodels.LibroModelForm{Titulo: &titulo, Sinopsis: &sipnosis, Creado: creado}

	nuevoLibro, err := librosmodels.CrearLibro(libro)
	if err != nil {
		t.Errorf("No se pudo crear el libro: %v", err)
		t.Fail()
	}
	oid = nuevoLibro.Hex()
	t.Logf("Libro creado con exito : %s", nuevoLibro)
}

func TestListarLibros(t *testing.T) {

	_, err := librosmodels.ListarLibros()

	if err != nil {
		t.Errorf("Error al obtener los libros: %v", err)
		t.Fail()
	}
}

func TestObtenerLibro(t *testing.T) {
	t.Run("libro existente", func(t *testing.T) {
		libroobtenido, err := librosmodels.ObtenerLibro("64a3b9bb60740a5a6707e64b")
		if err != nil {
			t.Fatalf("Error al obtenr el libro: %v", err)
			t.Fail()
		}
		t.Logf("libro obtenido: %v", libroobtenido)

	})
	t.Run("libro no existe", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != nil {
				t.Error(err)
				t.Fail()
			}
		}()

		_, err := librosmodels.ObtenerLibro("64a3b9bb60740a5a6707e647")
		if err != nil {
			t.Logf("Error al obtenr el libro: %v", err)
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

		_, err := librosmodels.ObtenerLibro("64a3b9bb60740a5a6707")
		if err != nil {
			t.Logf("Error al obtenr el libro: %v", err)
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
		t.Errorf("Error al actualizar libro: %v", err)
		t.Fail()

	}
}

func TestEliminarLibro(t *testing.T) {
	/* err := librosmodels.EliminarLibro(oid)
	if err != nil {
		t.Errorf("Error al eliminar libro: %v", err)
		t.Fail()

	} */
}
