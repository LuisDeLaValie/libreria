package autormodels

import (
	"context"
	"time"

	"github.com/TDTxLE/libreria/database"
	"github.com/TDTxLE/libreria/models"
	"github.com/TDTxLE/libreria/utils"
	utilsmongo "github.com/TDTxLE/libreria/utils/mongoTogolang"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AutorModel struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nombre      string             `json:"nombre,omitempty" bson:"nombre,omitempty"`
	Creado      *time.Time         `json:"creado,omitempty" bson:"creado,omitempty"`
	Actualizado *time.Time         `json:"actualizado,omitempty" bson:"actualizado,omitempty"`
}

const dbCollection = "autores"

func CrearAutor(nuevoLibro AutorModel) (*primitive.ObjectID, error) {

	defer func() {
		database.Desconectar()
	}()

	database.Conectar()

	// Insertar el documento en la colección
	oid, err := database.Collection(dbCollection).InsertOne(context.TODO(), nuevoLibro)
	if err != nil {
		// Manejar el error de inserción
		database.Desconectar()
		statuscode := utils.GetHTTPStatusCode(err)
		return nil, models.ResposeError{
			Status:     "No se puedo crear el Autor",
			StatusCode: &statuscode,
			Message:    "No se pudo crear eldocumento",
			Detalle:    err,
		}
	}

	id := oid.InsertedID.(primitive.ObjectID)

	database.Desconectar()
	return &id, nil
}

func ListarAutores() ([]AutorModel, error) {

	defer func() {
		database.Desconectar()
	}()

	var autores []AutorModel
	database.Conectar()
	// Realiza una consulta para obtener múltiples documentos
	cursor, err := database.Collection(dbCollection).Find(context.Background(), bson.M{})
	if err != nil {
		database.Desconectar()
		statuscode := utils.GetHTTPStatusCode(err)
		return nil, models.ResposeError{
			Status:     "No se puedo crear el Autor",
			StatusCode: &statuscode,
			Message:    "Error al leer el documentos",
			Detalle:    err,
		}
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var autor AutorModel
		// Decodifica el documento en la variable result
		err := cursor.Decode(&autor)
		if err != nil {
			database.Desconectar()
			statuscode := utils.GetHTTPStatusCode(err)
			return nil, models.ResposeError{
				Status:     "No se puedo crear el Autor",
				StatusCode: &statuscode,
				Message:    "Error al leer el documentos",
				Detalle:    err,
			}
		}
		autores = append(autores, autor)
	}

	database.Desconectar()
	return autores, nil
}

func ObtenerAutor(oid string) (*AutorModel, error) {

	defer func() {
		database.Desconectar()
	}()

	database.Conectar()

	// Crear un ID de tipo ObjectID a partir de una cadena
	id, err := utilsmongo.ParseOID(oid)
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
	var result AutorModel
	err = database.Collection(dbCollection).FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		/* if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("No se encontró el documento: %v", err)
		} else {
			panic(err)
		} */

		statuscode := utils.GetHTTPStatusCode(err)
		return nil, models.ResposeError{
			Status:     "no se encontro autor",
			StatusCode: &statuscode,
			Message:    "No se encontró el documento",
			Detalle:    err,
		}
	}

	database.Desconectar()

	return &result, nil
}

func ActualizarAutor(oid string, actualizar AutorModel) error {
	defer func() {
		database.Desconectar()
	}()

	database.Conectar()

	// Crear un ID de tipo ObjectID a partir de una cadena
	id, err := utilsmongo.ParseOID(oid)
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
	current := time.Now()
	actualizar.Actualizado = &current
	update := bson.M{"$set": actualizar}

	// Realizar la actualización del documento
	_, err = database.Collection(dbCollection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		statuscode := utils.GetHTTPStatusCode(err)
		return models.ResposeError{
			Status:     "no se encontro autor",
			StatusCode: &statuscode,
			Message:    "No se encontró el documento",
			Detalle:    err,
		}
	}

	database.Desconectar()

	return nil
}

func EliminarAutor(oid string) error {
	defer func() {
		database.Desconectar()
	}()

	database.Conectar()

	// Crear un ID de tipo ObjectID a partir de una cadena
	id, err := utilsmongo.ParseOID(oid)
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
			Status:     "Eliminar archivo",
			StatusCode: &statuscode,
			Message:    "Error al obtener el id",
			Detalle:    err,
		}
	}

	database.Desconectar()

	return nil
}
