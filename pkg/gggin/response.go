package gggin

type Response[T any] struct {
	Data *T `json:"data"`
}

func NewResponse[T any](data T) *Response[T] {
	return &Response[T]{Data: &data}
}

var Ok = NewResponse("ok")
