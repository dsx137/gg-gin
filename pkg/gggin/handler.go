package gggin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler[T any] func(c *gin.Context) (*Response[T], *HttpError)

func ToGinHandler[T any](handler Handler[T]) func(c *gin.Context) {
	return func(c *gin.Context) {
		res, err := handler(c)
		if err != nil {
			c.JSON(err.StatusCode, NewResponse(err.Message))
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
