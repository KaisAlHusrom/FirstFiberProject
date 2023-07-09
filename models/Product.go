package models

import (
	"time"
)

type Product struct {
	Id          int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREAMENT;primary key"`
	ProductName string    `json:"product_name"`
	Price       int       `json:"price"`
	CreateAt    time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserId      int       `json:"user_id"`
}
