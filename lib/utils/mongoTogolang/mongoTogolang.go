package utilsmongo

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// cambierte el dato fecha de mongo al dato fecha de go
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

func GoStructToBson(gostruct interface{}) bson.M {
	auxdata := bson.M{}

	// Obtener el tipo del struct
	keys := reflect.TypeOf(gostruct)
	values := reflect.ValueOf(gostruct)

	// Iterar sobre los campos del struct
	for i := 0; i < keys.NumField(); i++ {
		campo := keys.Field(i)
		nombre := campo.Name

		valor := values.FieldByName(nombre)
		if !valor.IsNil() {

			if valor.Kind() == reflect.Struct {
				auxdata[strings.ToLower(nombre)] = GoStructToBson(valor.Elem().Interface())
			} else {
				auxdata[strings.ToLower(nombre)] = valor.Elem().Interface()
			}
		}

	}
	return auxdata
}

func ParseOID(oid string) (*primitive.ObjectID, error) {
	// Crear un ID de tipo ObjectID a partir de una cadena
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return nil, fmt.Errorf("oid incorecto: %s", err.Error())
	}
	return &id, nil
}

func ParseManyOID(oids []string) (*[]primitive.ObjectID, error) {
	// Slice de ObjectID para almacenar los valores convertidos
	objectIDs := make([]primitive.ObjectID, len(oids))

	// Convertir cada string a ObjectID
	for i, str := range oids {
		oid, err := primitive.ObjectIDFromHex(str)
		if err != nil {
			// Manejar el error en caso de que la conversión falle
			fmt.Println("Error:", err)
			return nil, fmt.Errorf("%s no es un oid valido : %s", str, err.Error())

		}
		objectIDs[i] = oid
	}

	return &objectIDs, nil
}
