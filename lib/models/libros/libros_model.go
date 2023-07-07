package models

import (
	"context"
	"fmt"
	"time"

	"github.com/TDTxLE/libreria/database"
	"github.com/TDTxLE/libreria/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func libromodelTobson(libro LibroModel) bson.M {
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
		libro.Creado = utils.MongotimeToGotime(auxcreado)
	}
	if data["actualizado"] != nil {
		auxactualizado := data["actualizado"].(primitive.DateTime)
		libro.Actualizado = utils.MongotimeToGotime(auxactualizado)
	}

	return libro
}

type LibroModel struct {
	Id          *string
	Titulo      *string
	Sinopsis    *string
	Origen      *origen
	Creado      *time.Time
	Actualizado *time.Time
}

type origen struct {
	Hosth string
	Url   string
}

func (lirbo LibroModel) CrearLibro(nuevoLibro LibroModel) (*LibroModel, error) {

	defer func() {
		database.Desconectar()
	}()

	database.Conectar()
	document := libromodelTobson(nuevoLibro)
	// Insertar el documento en la colección
	oid, err := database.Collection("libros").InsertOne(context.TODO(), document)
	if err != nil {
		// Manejar el error de inserción
		database.Desconectar()
		return nil, fmt.Errorf("No se pudo crear eldocumento: %v", err)
	}

	insertedID := oid.InsertedID.(primitive.ObjectID).Hex()
	nuevoLibro.Id = &insertedID

	database.Desconectar()
	return &nuevoLibro, nil
}

func (lirbo LibroModel) ListarLibros() ([]LibroModel, error) {

	defer func() {
		database.Desconectar()
	}()

	var libros []LibroModel
	database.Conectar()
	// Realiza una consulta para obtener múltiples documentos
	cursor, err := database.Collection("libros").Find(context.Background(), bson.M{})
	if err != nil {
		database.Desconectar()
		return nil, fmt.Errorf("Error al leer el documentos: %v", err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {

		// Declara una variable para almacenar el resultado de la consulta
		var result bson.M
		// Decodifica el documento en la variable result
		err := cursor.Decode(&result)
		if err != nil {
			database.Desconectar()
			// Manejar el error de decodificación
			panic(err)
		}
		newlibro := bsonTolibromodel(result)
		libros = append(libros, newlibro)

	}

	database.Desconectar()
	return libros, nil
}

func (lirbo LibroModel) ObtenerLibro(oid string) (*LibroModel, error) {

	defer func() {
		database.Desconectar()
	}()

	database.Conectar()

	// Crear un ID de tipo ObjectID a partir de una cadena
	id, err := utils.ValidarOID(oid)
	if err != nil {
		return nil, err
	}

	// Crear una consulta (query) para obtener un documento por ID
	filter := bson.M{"_id": id}

	// Obtener el documento correspondiente a la consulta
	var result bson.M
	err = database.Collection("libros").FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("No se encontró el documento: %v", err)
		} else {
			panic(err)
		}
	}

	database.Desconectar()
	libroObtenido := bsonTolibromodel(result)

	return &libroObtenido, nil
}

func (lirbo LibroModel) ActualizarLibro(oid string, actualizar LibroModel) error {
	defer func() {
		database.Desconectar()
	}()

	database.Conectar()

	// Crear un ID de tipo ObjectID a partir de una cadena
	id, err := utils.ValidarOID(oid)
	if err != nil {
		return err
	}

	// Crear una consulta (query) para obtener un documento por ID
	filter := bson.M{"_id": id}

	// Crear una actualización con los cambios deseados
	current := time.Now()
	actualizar.Actualizado = &current
	newdata := libromodelTobson(actualizar)
	update := bson.M{"$set": newdata}

	// Realizar la actualización del documento
	_, err = database.Collection("libros").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	database.Desconectar()

	return nil
}
func (lirbo LibroModel) EliminarLibro(oid string) error {
	defer func() {
		database.Desconectar()
	}()

	database.Conectar()

	// Crear un ID de tipo ObjectID a partir de una cadena
	id, err := utils.ValidarOID(oid)
	if err != nil {
		return err
	}

	// Crear una consulta (query) para obtener un documento por ID
	filter := bson.M{"_id": id}

	// Realizar la actualización del documento
	_, err = database.Collection("libros").DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil
	}

	database.Desconectar()

	return nil
}
