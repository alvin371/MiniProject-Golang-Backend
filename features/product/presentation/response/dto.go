package response

import (
	"A-Golang-MiniProject/features/product"
	"time"
)

type Product struct {
	ID          int                   `json:"id"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
	Name        string                `json:"name"`
	Price       float32               `json:"price"`
	Satuan      int                   `json:"satuan"`
	Distributor []product.Distributor `json:"Distributor"`
}

func FromCore(core product.Core) Product {
	return Product{
		ID:        core.ID,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
		Name:      core.Name,
		Price:     core.Price,
		Satuan:    core.Satuan,
	}
}

func FromCoreSlice(core []product.Core) []Product {
	var artArray []Product
	for key := range core {
		artArray = append(artArray, FromCore(core[key]))
	}
	return artArray
}
