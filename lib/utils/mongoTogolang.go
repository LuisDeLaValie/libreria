package utils

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidarOID(oid string) (*primitive.ObjectID, error) {
	if len(oid) != 24 {
		return nil, fmt.Errorf("oid incorecto")
	}

	// Crear un ID de tipo ObjectID a partir de una cadena
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		panic(err)
	}
	return &id, nil
}

func MongotimeToGotime(eltime primitive.DateTime) *time.Time {

	newtime := time.Unix(0, int64(eltime)*int64(time.Millisecond))
	return &newtime
}

func BsonToGoValue(data bson.M, key string) *interface{} {

	value, ok := data[key]
	if !ok {
		return nil
	}
	return &value
}
