package models

import (
	"fmt"
)

type ResposeError struct {
	Status     string `json:"status,omitempty"`
	StatusCode *int   `json:"-,omitempty"`
	Message    string `json:"mensaje"`
	Detalle    error  `json:"detalle,omitempty"`
}

func (e ResposeError) Error() string {
	return fmt.Sprintf(`{"status":"%s", "mensaje":"%s", "detalle":{%s}}"`, e.Status, e.Message, e.Detalle.Error())

}
