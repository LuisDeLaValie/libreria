package models

import "encoding/json"

type ResposeError struct {
	Status     string      `json:"status,omitempty"`
	StatusCode *int        `json:"-,omitempty"`
	Message    string      `json:"mensaje"`
	Detalle    interface{} `json:"detalle,omitempty"`
}

func (e ResposeError) Error() string {
	jsonData, err := json.Marshal(e)
	if err != nil {
		return e.Message
	}
	return string(jsonData)
}
