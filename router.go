package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/probe", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})
	return r
}
