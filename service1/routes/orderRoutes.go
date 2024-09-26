package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oguzhankuzlukluoglu/Basic-Service/service1/controllers"
)

func Router(app *gin.Engine) {
	r := app.Group("/myapi/v1")

	r.POST("createOrder", controllers.CreateOrder)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello"})
	})
}
