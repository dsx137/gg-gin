package gggin

import "github.com/gin-gonic/gin"

type Nothing any

func Set[T any](c *gin.Context, key string, value T) {
	c.Set(key, value)
}

func Get[T any](c *gin.Context, key string) (T, bool) {
	ret, ok := c.Get(key)
	if !ok {
		return *new(T), false
	}
	return ret.(T), true
}
