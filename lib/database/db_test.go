package database_test

import (
	"testing"

	"github.com/TDTxLE/libreria/database"
	"github.com/TDTxLE/libreria/utils"
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
				t.Error("probanod volumen")
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
