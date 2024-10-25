package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":        200,
			"language":    "go version go1.20.6",
			"message":     "Server running on port 8000",
			"description": "Belajar docker container",
			"version":     "v1.0",
		})
	})
	r.Run("0.0.0.0:8000")
}
