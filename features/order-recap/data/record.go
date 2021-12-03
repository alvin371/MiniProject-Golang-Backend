package data

import (
	"A-Golang-MiniProject/features/order-recap"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID          int
	Total_Price int
	Worker      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func toDomain(or Order) order.Core {
	return order.Core{
		ID:          or.ID,
		Total_Price: or.Total_Price,
		Worker:      or.Worker,
		CreatedAt:   or.CreatedAt,
		UpdatedAt:   or.UpdatedAt,
	}
}

func fromDomain(order order.Core) Order {
	return Order{
		ID:          order.ID,
		Total_Price: order.Total_Price,
		Worker:      order.Worker,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}
}

func toDomainList(data []Order) []order.Core {
	result := []order.Core{}

	for _, or := range data {
		result = append(result, toDomain(or))
	}
	return result
}
