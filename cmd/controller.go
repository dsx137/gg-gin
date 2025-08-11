package main

import (
	"github.com/dsx137/gg-gin/pkg/gggin"
	"github.com/gin-gonic/gin"
)

type ControllerIndex struct{}

func NewControllerIndex(g *gin.RouterGroup) *ControllerIndex {
	ctl := &ControllerIndex{}
	g.GET("/", gggin.HandleController(ctl.Handle))
	return ctl
}

func (ctl *ControllerIndex) Handle(c *gin.Context) (*gggin.Response[string], *gggin.HttpError) {
	return gggin.NewResponse("hello, world"), nil
}
