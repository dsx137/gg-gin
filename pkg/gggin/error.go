package gggin

import (
	"errors"

	"github.com/dsx137/gg-kit/pkg/ggkit"
)

type HttpError struct {
	StatusCode int
	Message    string
}

func NewHttpError(statusCode int, message string) *HttpError {
	return &HttpError{StatusCode: statusCode, Message: message}
}

func (e *HttpError) Error() string { return e.Message }

// EXPERIMENTAL: 不建议在统一流程控制错误处理中使用
func AsHttpErrorOrElse(err error, defaultCode int) *HttpError {
	if err == nil {
		return nil
	}

	httpErr, ok := ggkit.BindPtrR[HttpError](func(x any) bool { return errors.As(err, x) })
	if ok {
		return httpErr
	}

	return NewHttpError(defaultCode, err.Error())
}
