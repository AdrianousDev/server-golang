package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Carrega os templates da pasta
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "documentation.html", nil)
	})

	router.Run() // listens on 0.0.0.0:8080 by default
}
