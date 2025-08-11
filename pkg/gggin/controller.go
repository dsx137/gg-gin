package gggin

import (
	"net/http"

	"github.com/gin-gonic/gin"

)

type Nothing any
type Controller[T any] func(c *gin.Context) (*Response[T], *HttpError)

func HandleController[T any](controller Controller[T]) func(c *gin.Context) {
	return func(c *gin.Context) {
		response, err := controller(c)
		if err != nil {
			c.JSON(err.StatusCode, NewResponse(err.Error()))
			return
		}
		c.JSON(http.StatusOK, response)
	}
}
