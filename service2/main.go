package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oguzhankuzlukluoglu/service2/config"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello"})
	})
	if err := r.Run(":8082"); err != nil {
		panic("Server :8082 failed to start")
	}
}
