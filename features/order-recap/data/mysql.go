package data

import (
	"A-Golang-MiniProject/features/order-recap"

	"gorm.io/gorm"
)

type MysqlOrderRepository struct {
	Conn *gorm.DB
}

func NewMySqlOrderRepository(conn *gorm.DB) order.Data {
	return &MysqlOrderRepository{
		Conn: conn,
	}
}

func (rep *MysqlOrderRepository) GetAll() ([]order.Core, error) {
	var ord []Order

	result := rep.Conn.Find(&ord)

	if result.Error != nil {
		return []order.Core{}, result.Error
	}
	return toDomainList(ord), nil
}
func (rep *MysqlOrderRepository) Create(domain *order.Core) (order.Core, error) {
	or := fromDomain(*domain)
	result := rep.Conn.Create(&or)

	if result.Error != nil {
		return order.Core{}, nil
	}

	return toDomain(or), nil
}
