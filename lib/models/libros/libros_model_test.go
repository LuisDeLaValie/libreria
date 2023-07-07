package models_test

import (
	"testing"
	"time"

	models "github.com/TDTxLE/libreria/models/libros"
)

var oid string

func TestCrearLibro(t *testing.T) {
	titulo := "Lirbo de prueba"
	sipnosis := "Este es un libro de prueba para ver el funcionaminto dela api"
	creado := time.Now()
	libro := models.LibroModel{Titulo: &titulo, Sinopsis: &sipnosis, Creado: &creado}

	nuevoLibro, err := libro.CrearLibro(libro)
	if err != nil {
		t.Errorf("No se pudo crear el libro: %v", err)
		t.Fail()
	}
	oid = *nuevoLibro.Id
	t.Logf("Libro creado con exito : %v", nuevoLibro)
}

func TestListarLibros(t *testing.T) {
	var lirbo models.LibroModel

	_, err := lirbo.ListarLibros()

	if err != nil {
		t.Errorf("Error al obtener los libros: %v", err)
		t.Fail()
	}
}

func TestObtenerLibro(t *testing.T) {
	var libro models.LibroModel
	t.Run("libro existente", func(t *testing.T) {
		libroobtenido, err := libro.ObtenerLibro("64a3b9bb60740a5a6707e64b")
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

		_, err := libro.ObtenerLibro("64a3b9bb60740a5a6707e647")
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

		_, err := libro.ObtenerLibro("64a3b9bb60740a5a6707")
		if err != nil {
			t.Logf("Error al obtenr el libro: %v", err)
		}

	})

}

func TestActualizarLibro(t *testing.T) {
	var libro models.LibroModel
	clock := time.Now()
	auxTitulo := "actualizar titulo " + clock.Local().GoString()
	libro.Titulo = &auxTitulo
	err := libro.ActualizarLibro("64a5cf610d0195d943d2ee10", libro)
	if err != nil {
		t.Errorf("Error al actualizar libro: %v", err)
		t.Fail()

	}
}

func TestEliminarLibro(t *testing.T) {
	var libro models.LibroModel
	err := libro.EliminarLibro(oid)
	if err != nil {
		t.Errorf("Error al eliminar libro: %v", err)
		t.Fail()

	}
}
