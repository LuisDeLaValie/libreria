package models_test

import (
	"encoding/json"
	"testing"

	models "github.com/TDTxLE/libreria/models"
)

func Test(t *testing.T) {

	// Datos JSON de ejemplo
	errorjson := `{
		"status":"",
		"statuss":0,
		"mensaje":"mensaje de prueba",
		"detalle":"mensaje de prueba"
	}`
	// Decodificar el JSON en una estructura
	var persona models.ResposeError
	err := json.Unmarshal([]byte(errorjson), &persona)
	if err != nil {
		panic(err)
	}
}
