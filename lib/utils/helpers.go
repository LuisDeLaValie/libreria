package utils

import (
	"net/http"
	"os"
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
