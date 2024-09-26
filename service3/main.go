package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oguzhankuzlukluoglu/service3/config"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	if err := r.Run(":8083"); err != nil {
		panic("Server :8083 failed to start")
	}
}
