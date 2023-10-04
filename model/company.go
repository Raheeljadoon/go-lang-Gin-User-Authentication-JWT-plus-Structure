package model

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      int
}
