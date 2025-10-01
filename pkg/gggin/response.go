package gggin

import "net/http"

type RawResponse[T any] struct {
	Data *T `json:"data"`
}

func NewRawResponse[T any](data T) *RawResponse[T] {
	return &RawResponse[T]{Data: &data}
}

type Response[T any] struct {
	StatusCode  int
	RawResponse *RawResponse[T]
}

func NewResponseWithStatusCode[T any](statusCode int, data T) *Response[T] {
	return &Response[T]{StatusCode: statusCode, RawResponse: NewRawResponse(data)}
}

func NewResponse[T any](data T) *Response[T] {
	return NewResponseWithStatusCode(http.StatusOK, data)
}

var Ok = NewResponse("ok")
