package models

import (
	"gorm.io/gorm"
	"time"
)

type Item struct {
    gorm.Model
    Id          uint      `gorm:"primary_key" json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Price       float64   `json:"price"`
    Status      string    `json:"status"`
    SellerId    uint      `json:"seller_id"`
    CreatedAt   time.Time `json:"created_at"`
}