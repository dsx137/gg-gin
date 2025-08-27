package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) { c.Status(http.StatusMethodNotAllowed) })
	r.NoRoute(func(c *gin.Context) { c.Status(http.StatusNotFound) })

	api := r.Group("/api")
	{
		NewControllerIndex(api.Group("/"))
	}
}

func main() {
	r := gin.Default()

	initRouter(r)

	addr := ":54738"

	r.Run(addr)
}
