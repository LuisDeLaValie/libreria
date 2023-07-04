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
		"mongodb://%s:%s@%s:%s/",
		// "mongodb://rootLibreriaTest:Comemierda@mongo:27017/Libreria_test?authSource=admin",
		utils.Getenv("DB_USER", "Biblioteca_User"),
		utils.Getenv("DB_PWD", "123456"),
		utils.Getenv("DB_HOST", "localhost"),
		utils.Getenv("DB_PORT", "12500"),
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

func Collection(coll string) *mongo.Collection {

	database := utils.Getenv("DB_DATABASE", "Libreria")

	return client.Database(database).Collection(coll)

}
