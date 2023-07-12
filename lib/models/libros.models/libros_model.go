package librosmodels

import (
	"context"
	"encoding/json"
	"time"

	"github.com/TDTxLE/libreria/database"
	"github.com/TDTxLE/libreria/models"
	autormodels "github.com/TDTxLE/libreria/models/autor.models"
	utils "github.com/TDTxLE/libreria/utils"
	utilsmongo "github.com/TDTxLE/libreria/utils/mongoTogolang"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LibroModelForm struct {
	Id          primitive.ObjectID    `json:"id,omitempty" bson:"_id,omitempty"`
	Titulo      *string               `json:"titulo,omitempty" bson:"titulo,omitempty"`
	Sinopsis    *string               `json:"sinopsis,omitempty" bson:"sinopsis,omitempty"`
	Origen      *Origen               `json:"origen,omitempty" bson:"origen,omitempty"`
	Autores     *[]primitive.ObjectID `json:"autores,omitempty" bson:"autores,omitempty"`
	Collection  primitive.ObjectID    `json:"collection,omitempty" bson:"collection,omitempty"`
	Creado      time.Time             `json:"creado,omitempty" bson:"creado,omitempty"`
	Actualizado time.Time             `json:"actualizado,omitempty" bson:"actualizado,omitempty"`
}
type LibroModel struct {
	Id          primitive.ObjectID        `json:"id,omitempty" bson:"_id,omitempty"`
	Titulo      *string                   `json:"titulo,omitempty" bson:"titulo,omitempty"`
	Sinopsis    *string                   `json:"sinopsis,omitempty" bson:"sinopsis,omitempty"`
	Origen      *Origen                   `json:"origen,omitempty" bson:"origen,omitempty"`
	Autores     *[]autormodels.AutorModel `json:"autores,omitempty" bson:"autores,omitempty"`
	Creado      *time.Time                `json:"creado,omitempty" bson:"creado,omitempty"`
	Actualizado *time.Time                `json:"actualizado,omitempty" bson:"actualizado,omitempty"`
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
	consulta := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "autores",
				"localField":   "autores",
				"foreignField": "_id",
				"as":           "autores",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":         1,
				"titulo":      1,
				"sinopsis":    1,
				"creado":      1,
				"actualizado": 1,
				"autores": bson.M{
					"$cond": bson.M{
						"if": bson.M{
							"$eq": bson.A{"$autores", bson.A{}},
						},
						"then": "$$REMOVE",
						"else": bson.M{
							"$map": bson.M{
								"input": "$autores",
								"as":    "autor",
								"in": bson.M{
									"_id":    "$$autor._id",
									"nombre": "$$autor.nombre",
								},
							},
						},
					},
				},
			},
		},
	}
	colecion := database.Collection(dbCollection)
	cursor, err := colecion.Aggregate(context.Background(), consulta)
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

	// Realiza una consulta para obtener múltiples documentos
	consulta := bson.A{
		bson.M{
			"$match": bson.M{
				"_id": id,
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "autores",
				"localField":   "autores",
				"foreignField": "_id",
				"as":           "autores",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":         1,
				"titulo":      1,
				"sinopsis":    1,
				"creado":      1,
				"actualizado": 1,
				"autores": bson.M{
					"$cond": bson.M{
						"if": bson.M{
							"$eq": bson.A{"$autores", bson.A{}},
						},
						"then": "$$REMOVE",
						"else": bson.M{
							"$map": bson.M{
								"input": "$autores",
								"as":    "autor",
								"in": bson.M{
									"_id":    "$$autor._id",
									"nombre": "$$autor.nombre",
								},
							},
						},
					},
				},
			},
		},
	}

	// Obtener el documento correspondiente a la consulta
	colecion := database.Collection(dbCollection)
	cursor, err := colecion.Aggregate(context.Background(), consulta)
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

	var result LibroModel
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
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

func ActualizarVariosLibros(oids []string, actualizar LibroModelForm) error {
	database.Conectar()
	idis := []primitive.ObjectID{}
	// Crear un ID de tipo ObjectID a partir de una cadena
	for i := 0; i < len(oids); i++ {
		id, err := utilsmongo.ValidarOID(oids[i])
		if err != nil {
			statuscode := utils.GetHTTPStatusCode(err)
			return models.ResposeError{
				Status:     "id no valid",
				StatusCode: &statuscode,
				Message:    "Error al obtener el id",
				Detalle:    err,
			}
		}
		idis = append(idis, *id)
	}

	// Crear una consulta (query) para obtener un documento por ID
	filter := bson.M{"_id": bson.M{"$in": idis}}

	// Crear una actualización con los cambios deseados
	actualizar.Actualizado = time.Now()
	update := bson.M{"$set": actualizar}
	_, err := database.Collection(dbCollection).UpdateMany(context.TODO(), filter, update)
	if err != nil {
		statuscode := utils.GetHTTPStatusCode(err)
		return models.ResposeError{
			Status:     "Error al actualizar los documentos",
			StatusCode: &statuscode,
			Message:    "no se pudo acutalizar los documentos",
			Detalle:    err,
		}
	}
	database.Desconectar()
	return nil
}

func EliminarVariosLibros(oids []string) error {
	database.Conectar()
	idis := []primitive.ObjectID{}
	// Crear un ID de tipo ObjectID a partir de una cadena

	for i := 0; i < len(oids); i++ {
		id, err := utilsmongo.ValidarOID(oids[i])
		if err != nil {
			statuscode := utils.GetHTTPStatusCode(err)
			return models.ResposeError{
				Status:     "id no valid",
				StatusCode: &statuscode,
				Message:    "Error al obtener el id",
				Detalle:    err,
			}
		}
		idis = append(idis, *id)
	}

	filter := bson.M{"_id": bson.M{"$in": idis}}
	_, err := database.Collection(dbCollection).DeleteMany(context.TODO(), filter)
	if err != nil {
		statuscode := utils.GetHTTPStatusCode(err)
		return models.ResposeError{
			Status:     "Error al Eliniar los documentos",
			StatusCode: &statuscode,
			Message:    "no se pudo Eliniar los documentos",
			Detalle:    err,
		}
	}

	database.Desconectar()
	return nil
}

func (a LibroModelForm) ToJson() string {
	jsonData, err := json.Marshal(a)
	if err != nil {
	}

	return string(jsonData)
}
