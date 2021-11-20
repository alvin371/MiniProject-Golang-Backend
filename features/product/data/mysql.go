package data

import (
	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	Conn *gorm.DB
}
