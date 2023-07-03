package database

import (
	"context"
	"fmt"

	"github.com/TDTxLE/libreria/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Conectar() {

	// Generar cadena de conexion mediante Environment
	uri := fmt.Sprintf(
		// "mongodb://%s:%s@%s:%s",
		"mongodb://%s:%s@%s:%s/?authSource=%s",
		utils.Getenv("DB_USER", "Biblioteca_User"),
		utils.Getenv("DB_PWD", "123456"),
		utils.Getenv("DB_HOST", "localhost"),
		utils.Getenv("DB_POST", "12500"),
		utils.Getenv("DB_DATABASE", "Libreria"),
	)

	// Configura las opciones de conexión
	clientOptions := options.Client().ApplyURI(uri)

	// Conecta al servidor MongoDB
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic("Error al conectar a MongoDB: " + err.Error())
	}

	// Comprueba la conexión
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic("Error al hacer ping a MongoDB: " + err.Error())
	}
}

func Desconectar() {
	// Desconecta del servidor MongoDB
	err := client.Disconnect(context.Background())
	if err != nil {
		panic("Error al desconectar de MongoDB: " + err.Error())
	}
}
