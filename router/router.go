package router

import "github.com/gin-gonic/gin"

func Initilizer() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	router.Run(":8080") // roda em http://localhost:8080
}