package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oguzhankuzlukluoglu/Basic-Service/service1/config"
	"github.com/oguzhankuzlukluoglu/Basic-Service/service1/routes"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	routes.Router(r)
	if err := r.Run(":8081"); err != nil {
		panic("Server :8081 failed to start")
	}
}
