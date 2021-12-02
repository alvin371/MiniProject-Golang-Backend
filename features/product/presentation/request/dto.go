package request

import (
	"A-Golang-MiniProject/features/product"
)

type Product struct {
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Satuan int     `json:"satuan"`
}

func ToCore(req Product) product.Core {
	return product.Core{
		Name:   req.Name,
		Price:  req.Price,
		Satuan: req.Satuan,
	}
}
