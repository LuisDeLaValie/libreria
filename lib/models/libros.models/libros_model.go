package librosmodels

import (
	"context"
	"time"

	"github.com/TDTxLE/libreria/database"
	"github.com/TDTxLE/libreria/models"
	autormodels "github.com/TDTxLE/libreria/models/autor.models"
	utils "github.com/TDTxLE/libreria/utils"
	utilsmongo "github.com/TDTxLE/libreria/utils/mongoTogolang"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* func libromodelTobson(libro LibroModel) bson.M {
	data := bson.M{}
	if libro.Titulo != nil {
		data["titulo"] = libro.Titulo
	}
	if libro.Sinopsis != nil {
		data["sinopsis"] = libro.Sinopsis
	}
	if libro.Creado != nil {
		data["creado"] = libro.Creado
	}
	if libro.Actualizado != nil {
		data["actualizado"] = libro.Actualizado
	}
	if libro.Origen != nil {
		data["origen"] = bson.M{
			"host": libro.Origen.Host,
			"url":  libro.Origen.Url,
		}

	}
	return data
}
func bsonTolibromodel(data bson.M) LibroModel {

	var libro LibroModel

	if data["_id"] != nil {
		auxoid := data["_id"].(primitive.ObjectID).Hex()
		libro.Id = &auxoid
	}
	if data["tituloo"] != nil {
		auxtitulo := data["tituloo"].(string)
		libro.Titulo = &auxtitulo
	}
	if data["sinopsis"] != nil {
		auxsinopsis := data["sinopsis"].(string)
		libro.Sinopsis = &auxsinopsis
	}
	if data["creado"] != nil {
		auxcreado := data["creado"].(primitive.DateTime)
		libro.Creado = utilsmongo.MongotimeToGotime(auxcreado)
	}
	if data["actualizado"] != nil {
		auxactualizado := data["actualizado"].(primitive.DateTime)
		libro.Actualizado = utilsmongo.MongotimeToGotime(auxactualizado)
	}

	return libro
} */

type LibroModelForm struct {
	Id          primitive.ObjectID    `json:"id,omitempty" bson:"_id,omitempty"`
	Titulo      *string               `json:"titulo,omitempty" bson:"titulo,omitempty"`
	Sinopsis    *string               `json:"sinopsis,omitempty" bson:"sinopsis,omitempty"`
	Origen      *Origen               `json:"origen,omitempty" bson:"origen,omitempty"`
	Autores     *[]primitive.ObjectID `json:"autores,omitempty" bson:"autores,omitempty"`
	Creado      time.Time             `json:"creado,omitempty" bson:"creado,omitempty"`
	Actualizado time.Time             `json:"actualizado,omitempty" bson:"actualizado,omitempty"`
}
type LibroModel struct {
	Id          primitive.ObjectID        `json:"id,omitempty" bson:"_id,omitempty"`
	Titulo      *string                   `json:"titulo,omitempty" bson:"titulo,omitempty"`
	Sinopsis    *string                   `json:"sinopsis,omitempty" bson:"sinopsis,omitempty"`
	Origen      *Origen                   `json:"origen,omitempty" bson:"origen,omitempty"`
	Autores     *[]autormodels.AutorModel `json:"autores,omitempty" bson:"autores,omitempty"`
	Creado      time.Time                 `json:"creado,omitempty" bson:"creado,omitempty"`
	Actualizado time.Time                 `json:"actualizado,omitempty" bson:"actualizado,omitempty"`
}

type Origen struct {
	Host string `json:"host" bson:"host"`
	Url  string `json:"url" bson:"url"`
}

const dbCollection = "libros"

func CrearLibro(nuevoLibro LibroModelForm) (*primitive.ObjectID, error) {

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
			Message:    "Error al leer el documentos",
			Detalle:    err,
		}
	}

	database.Desconectar()
	id := oid.InsertedID.(primitive.ObjectID)
	return &id, nil
}

func ListarLibros() ([]LibroModel, error) {

	defer func() {
		database.Desconectar()
	}()

	var libros []LibroModel
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
		var result LibroModel
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

func ObtenerLibro(oid string) (*LibroModel, error) {

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
	var result LibroModel
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

func ActualizarLibro(oid string, actualizar LibroModelForm) error {
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
func EliminarLibro(oid string) error {
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
