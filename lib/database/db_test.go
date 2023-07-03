package database_test

import (
	"testing"

	"github.com/TDTxLE/libreria/database"
)

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
