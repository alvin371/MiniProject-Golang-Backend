package request

import "A-Golang-MiniProject/features/order-recap"

type Order struct {
	Total_Price int    `json:"total_price"`
	Worker      string `json:"worker_name"`
}

func (req *Order) ToDomain() order.Core {
	return order.Core{
		Total_Price: req.Total_Price,
		Worker:      req.Worker,
	}
}
