package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oguzhankuzlukluoglu/Basic-Service/service1/config"
	"github.com/oguzhankuzlukluoglu/Basic-Service/service1/routes"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello"})
	})
	routes.Router(r)
	if err := r.Run(":8081"); err != nil {
		panic("Server failed to start")
	}
}
