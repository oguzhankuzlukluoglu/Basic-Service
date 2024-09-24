package models

import (
	"time"
)

type Order struct {
	OrderID uint      `json:"order_id" gorm:"primaryKey;autoIncrement"`
	UserID  uint      `json:"user_id"`
	Total   float64   `json:"total"`
	Date    time.Time `json:"date"`
	IsPaid  bool      `json:"is_paid" gorm:"default:false"`
}
