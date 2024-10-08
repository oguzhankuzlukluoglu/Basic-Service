package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oguzhankuzlukluoglu/Basic-Service/service1/controllers"
)

func Router(app *gin.Engine) {
	r := app.Group("/myapi/v1")

	r.POST("createOrder", controllers.CreateOrder)
	r.POST("deneme", controllers.CreateOrder)
	r.GET("/get", controllers.GetMessage)
}
