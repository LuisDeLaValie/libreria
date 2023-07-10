package collecionmodels

import (
	"context"
	"time"

	utilsmongo "github.com/TDTxLE/libreria/utils/mongoTogolang"

	"github.com/TDTxLE/libreria/database"
	"github.com/TDTxLE/libreria/models"
	"github.com/TDTxLE/libreria/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ColleccionModel struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Titulo      *string            `json:"titulo,omitempty" bson:"titulo,omitempty"`
	Sinopsis    *string            `json:"sinopsis,omitempty" bson:"sinopsis,omitempty"`
	Creado      time.Time          `json:"creado,omitempty" bson:"creado,omitempty"`
	Actualizado time.Time          `json:"actualizado,omitempty" bson:"actualizado,omitempty"`
}

const dbCollection = "colleciones"

func CrearColeccion(nuevoColeccion ColleccionModel) (*primitive.ObjectID, error) {

	defer func() {
		database.Desconectar()
	}()

	database.Conectar()

	// Insertar el documento en la colección
	oid, err := database.Collection(dbCollection).InsertOne(context.TODO(), nuevoColeccion)
	if err != nil {
		// Manejar el error de inserción
		database.Desconectar()
		statuscode := utils.GetHTTPStatusCode(err)
		return nil, models.ResposeError{
			Status:     "No se puedo crear el Autor",
			StatusCode: &statuscode,
			Message:    "Error al leer el documentos",
			Detalle:    err,
		}
	}

	database.Desconectar()
	id := oid.InsertedID.(primitive.ObjectID)
	return &id, nil
}

func ListarColecciones() ([]ColleccionModel, error) {

	defer func() {
		database.Desconectar()
	}()

	var libros []ColleccionModel
	database.Conectar()
	// Realiza una consulta para obtener múltiples documentos
	cursor, err := database.Collection(dbCollection).Find(context.Background(), bson.M{})
	if err != nil {
		database.Desconectar()
		statuscode := utils.GetHTTPStatusCode(err)
		return nil, models.ResposeError{
			Status:     "no se pudo obtener libros",
			StatusCode: &statuscode,
			Message:    "Error al leer el documentos",
			Detalle:    err,
		}
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {

		// Declara una variable para almacenar el resultado de la consulta
		var result ColleccionModel
		// Decodifica el documento en la variable result
		err := cursor.Decode(&result)
		if err != nil {
			database.Desconectar()
			statuscode := utils.GetHTTPStatusCode(err)
			return nil, models.ResposeError{
				Status:     "Error al comvertir los datos",
				StatusCode: &statuscode,
				Message:    "Error al leer el documentos",
				Detalle:    err,
			}
		}
		libros = append(libros, result)

	}

	database.Desconectar()
	return libros, nil
}

func ObtenerColeccion(oid string) (*ColleccionModel, error) {

	defer func() {
		database.Desconectar()
	}()

	database.Conectar()

	// Crear un ID de tipo ObjectID a partir de una cadena
	id, err := utilsmongo.ValidarOID(oid)
	if err != nil {
		statuscode := utils.GetHTTPStatusCode(err)
		return nil, models.ResposeError{
			Status:     "id no valid",
			StatusCode: &statuscode,
			Message:    "Error al obtener el id",
			Detalle:    err,
		}
	}

	// Crear una consulta (query) para obtener un documento por ID
	filter := bson.M{"_id": id}

	// Obtener el documento correspondiente a la consulta
	var result ColleccionModel
	err = database.Collection(dbCollection).FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		/* if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("No se encontró el documento: %v", err)
		} else {
			panic(err)
		} */

		statuscode := utils.GetHTTPStatusCode(err)
		return nil, models.ResposeError{
			Status:     "Error al leer el documento",
			StatusCode: &statuscode,
			Message:    "no se obtubo el documento",
			Detalle:    err,
		}
	}

	database.Desconectar()

	return &result, nil
}

func ActualizarColeccion(oid string, actualizar ColleccionModel) error {
	defer func() {
		database.Desconectar()
	}()

	database.Conectar()

	// Crear un ID de tipo ObjectID a partir de una cadena
	id, err := utilsmongo.ValidarOID(oid)
	if err != nil {
		statuscode := utils.GetHTTPStatusCode(err)
		return models.ResposeError{
			Status:     "id no valid",
			StatusCode: &statuscode,
			Message:    "Error al obtener el id",
			Detalle:    err,
		}
	}

	// Crear una consulta (query) para obtener un documento por ID
	filter := bson.M{"_id": id}

	// Crear una actualización con los cambios deseados
	actualizar.Actualizado = time.Now()
	update := bson.M{"$set": actualizar}

	// Realizar la actualización del documento
	_, err = database.Collection(dbCollection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		statuscode := utils.GetHTTPStatusCode(err)
		return models.ResposeError{
			Status:     "Error al leer el documento",
			StatusCode: &statuscode,
			Message:    "no se obtubo el documento",
			Detalle:    err,
		}
	}

	database.Desconectar()

	return nil
}

func EliminarColeccion(oid string) error {
	defer func() {
		database.Desconectar()
	}()

	database.Conectar()

	// Crear un ID de tipo ObjectID a partir de una cadena
	id, err := utilsmongo.ValidarOID(oid)
	if err != nil {
		statuscode := utils.GetHTTPStatusCode(err)
		return models.ResposeError{
			Status:     "id no valid",
			StatusCode: &statuscode,
			Message:    "Error al obtener el id",
			Detalle:    err,
		}
	}

	// Crear una consulta (query) para obtener un documento por ID
	filter := bson.M{"_id": id}

	// Realizar la actualización del documento
	_, err = database.Collection(dbCollection).DeleteOne(context.TODO(), filter)
	if err != nil {
		statuscode := utils.GetHTTPStatusCode(err)
		return models.ResposeError{
			Status:     "Error al eliminar el documento",
			StatusCode: &statuscode,
			Message:    "no se pudo borrar el documento",
			Detalle:    err,
		}
	}

	database.Desconectar()

	return nil
}
