package main

import (
	"net/http"

	"github.com/dsx137/gg-gin/pkg/gggin"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gggin.NewResponse("Method Not Allowed"))
	})

	r.NoRoute(func(c *gin.Context) { c.Status(404) })

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
