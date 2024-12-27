package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
    gorm.Model
    Id          	uint      `gorm:"primary_key" json:"id"`
    ItemId        	string    `json:"name"`
    BuyerId 		uint      `json:"description"`
    TransactionDate time.Time `json:"price"`
    Status      	string    `json:"status"`
    CreatedAt   	time.Time `json:"created_at"`
}