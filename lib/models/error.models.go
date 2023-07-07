package models

type ResponseError struct {
	Status     string `json:"status"`
	StatusCode *int   `json:"-"`
	Message    string `json:"mensaje"`

	Detalle *any `json:"detalle"`
}

func (e *ResponseError) Error() string {
	return e.Message
}
