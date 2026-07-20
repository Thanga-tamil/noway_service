package router

import (
	"github.com/gin-gonic/gin"
	"gateway/internal/api/rest/handler"
)

func Route(v1 *gin.RouterGroup)  {

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1.POST("/register", handler.HandleUserRegister)
	v1.GET("/jwt", handler.GenerateJwtToken)

}

