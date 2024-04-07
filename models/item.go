package models

import "gorm.io/gorm"

// gorm.ModelにはID, CreatedAt, UpdatedAt, DeletedAtが含まれるため、もともとあったIDを削除
type Item struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Price       uint   `gorm:"not null"`
	Description string
	SoldOut     bool `gorm:"not null;default:false"`
}
