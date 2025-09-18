package router

import "github.com/gin-gonic/gin"

func InitilizerRoutes(router *gin.Engine) {
router.Group( "/api/v1") 
}