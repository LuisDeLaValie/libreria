package database_test

import (
	"context"
	"testing"

	"github.com/TDTxLE/libreria/database"
	"github.com/TDTxLE/libreria/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func TestProvarEnv(t *testing.T) {

	DB_USER := utils.Getenv("DB_USER", "pato")
	if DB_USER == "pato" {
		t.Error("la ENV DB_USER no es lo que se espera")
		t.Fail()
	} else {
		t.Log("DB_USER correcto")
	}
	DB_PWD := utils.Getenv("DB_PWD", "pato")
	if DB_PWD == "pato" {
		t.Error("la ENV DB_PWD  no es lo que se espera")
		t.Fail()
	} else {
		t.Log("DB_PWD correcto")
	}
	DB_HOST := utils.Getenv("DB_HOST", "pato")
	if DB_HOST == "pato" {
		t.Error("la ENV DB_HOST no es lo que se espera")
		t.Fail()
	} else {
		t.Log("DB_HOST correcto")
	}
	DB_PORT := utils.Getenv("DB_PORT", "pato")
	if DB_PORT == "pato" {
		t.Error("la ENV DB_PORT no es lo que se espera")
		t.Fail()
	} else {
		t.Log("DB_PORT correcto")
	}
	DB_DATABASE := utils.Getenv("DB_DATABASE", "pato")
	if DB_DATABASE == "pato" {
		t.Error("la ENV DB_DATA no es lo que se espera")
		t.Fail()
	} else {
		t.Log("DB_DATABASE correcto")
	}

}

func TestConexionDB(t *testing.T) {

	t.Run("Probar metodo Conectar", func(t *testing.T) {

		defer func() {
			err := recover()
			if err != nil {
				t.Error(err)
				t.Fail()
			}
		}()

		database.Conectar()
	})
	t.Run("Probar metodo Desconectar", func(t *testing.T) {

		defer func() {
			err := recover()
			if err != nil {
				t.Error(err)
				t.Fail()
			}
		}()

		database.Conectar()
	})
}

func TestProbandoConsultas(t *testing.T) {

	t.Run("crear documento", func(t *testing.T) {

		defer func() {
			err := recover()
			if err != nil {
				t.Error(err)
				t.Fail()
			}
		}()
		// Crea un documento que deseas insertar
		document := bson.M{"name": "John Doe", "age": 30}

		database.Conectar()

		// Inserta el documento en la colección
		_, err := database.Collection("probar_test").InsertOne(context.Background(), document)
		if err != nil {
			t.Error("Error al insertar el documento: " + err.Error())
			t.Fail()
		} else {
			t.Log("se creo el documento")
		}
		database.Desconectar()
	})

	t.Run("leer datos", func(t *testing.T) {

		defer func() {
			err := recover()
			if err != nil {
				t.Error(err)
				t.Fail()
			}
		}()
		database.Conectar()
		// Realiza una consulta para obtener múltiples documentos
		cursor, err := database.Collection("probar_test").Find(context.Background(), bson.M{})
		if err != nil {
			t.Errorf("Error al leer los documentos: %v", err)
			t.Fail()
		}
		defer cursor.Close(context.Background())
		cont := 0
		for cursor.Next(context.Background()) {
			cont++
		}

		if cont == 0 {
			t.Error("Nose pudo obtener los datos")
			t.Fail()
		} else {
			t.Logf("Documentos leídos exitosamente :%d", cont)
		}

		database.Desconectar()
	})

}
