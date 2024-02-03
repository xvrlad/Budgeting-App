package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name string
	Price float32
	WantDate time.Time
}