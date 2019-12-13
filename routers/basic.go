package routers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) {
	router.POST("/test", GoPool)
}

func GoPool(c *gin.Context) {
	log.Printf("this is method")
}
