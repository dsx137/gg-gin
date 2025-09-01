package gggin

import (
	"github.com/dsx137/gg-kit/pkg/ggkit"
	"github.com/gin-gonic/gin"
)

func ShouldBindJSON[T any](c *gin.Context) (*T, error)   { return ggkit.BindR[T](c.ShouldBindJSON) }
func ShouldBindQuery[T any](c *gin.Context) (*T, error)  { return ggkit.BindR[T](c.ShouldBindQuery) }
func ShouldBindUri[T any](c *gin.Context) (*T, error)    { return ggkit.BindR[T](c.ShouldBindUri) }
func ShouldBindHeader[T any](c *gin.Context) (*T, error) { return ggkit.BindR[T](c.ShouldBindHeader) }
func ShouldBindXML[T any](c *gin.Context) (*T, error)    { return ggkit.BindR[T](c.ShouldBindXML) }
func ShouldBindYAML[T any](c *gin.Context) (*T, error)   { return ggkit.BindR[T](c.ShouldBindYAML) }
func ShouldBindTOML[T any](c *gin.Context) (*T, error)   { return ggkit.BindR[T](c.ShouldBindTOML) }
