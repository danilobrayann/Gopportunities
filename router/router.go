package router

import "github.com/gin-gonic/gin"

func Initilizer() {
	router  := gin.Default()

      // inicilaze routes

	 
	// run the serve
	router.Run(":8080") // roda em http://localhost:8080
}