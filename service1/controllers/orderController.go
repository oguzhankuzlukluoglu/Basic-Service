package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oguzhankuzlukluoglu/Basic-Service/service1/config"
	"github.com/oguzhankuzlukluoglu/Basic-Service/service1/models"
)

func CreateOrder(c *gin.Context) {
	var orderData models.Order

	if err := c.BindJSON(&orderData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrder := models.Order{
		UserID: orderData.UserID,
		Total:  orderData.Total,
		Date:   time.Now(),
	}

	result := config.Db.Create(&newOrder)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, newOrder)
}
