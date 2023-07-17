package utils

import (
	"net/http"
	"net/url"
	"os"

	utilsmongo "github.com/TDTxLE/libreria/utils/mongoTogolang"
	"go.mongodb.org/mongo-driver/bson"
)

func Getenv(key, defaultValue string) string {
	value, defined := os.LookupEnv(key)

	if !defined {
		return defaultValue
	}

	return value
}

func GetHTTPStatusCode(err error) int {
	// Verificar si el error implementa la interfaz HTTPError
	if httpErr, ok := err.(interface{ HTTPStatus() int }); ok {
		return httpErr.HTTPStatus()
	}

	// Por defecto, devolver código 500 Internal Server Error
	return http.StatusInternalServerError
}

func QueryParamsToBson(queryParams url.Values) *bson.M {
	if len(queryParams) > 0 {
		// Convertir los query parameters a un filtro BSON
		filter := bson.M{}
		for key, values := range queryParams {
			// cambar el query id por _id
			var auxkey string
			if key == "id" {
				auxkey = "_id"
			} else {
				auxkey = key
			}

			if len(values) == 1 {
				n, err := utilsmongo.ParseOID(values[0])
				if err != nil {
					filter[auxkey] = values[0]
				} else {
					filter[auxkey] = n
				}
			} else {
				n, err := utilsmongo.ParseManyOID(values)
				if err != nil {
					filter[auxkey] = bson.M{"$in": values}
				} else {
					filter[auxkey] = bson.M{"$in": n}
				}

			}
		}
		return &filter
	} else {
		return nil
	}

}
