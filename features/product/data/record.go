package data

import (
	"A-Golang-MiniProject/features/product"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50)"`
	Price       float32
	Satuan      int
	Distributor []Distributor `gorm:"many2many:product_distributor"`
}

type Distributor struct {
	gorm.Model
	Name    string `gorm:"type:varchar(50)"`
	Telp    int
	Address string `gorm:"type:varchar(225)"`
}

// DTO
func DistributorRecord(distributor []product.Distributor) []Distributor {
	convertedDistributor := []Distributor{}
	for _, s := range distributor {
		convertedDistributor = append(convertedDistributor, Distributor{
			Name: s.Name,
		})
	}

	return convertedDistributor
}

func ProductRecord(products product.Core) Product {
	return Product{
		Name:        products.Name,
		Price:       products.Price,
		Satuan:      products.Satuan,
		Distributor: DistributorRecord(products.Distributor),
	}
}
