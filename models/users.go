package models

import (
	"github.com/jinzhu/gorm"
)

// User is ...
type User struct {
	gorm.Model
	Name     string `json:"title"`
	LastName string `json:"author"`
	DNI      string `json:"dni"`
}
