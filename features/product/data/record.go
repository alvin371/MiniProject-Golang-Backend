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

func (p *Product) toCore() product.Core {
	return product.Core{
		ID:        int(p.ID),
		Name:      p.Name,
		Price:     p.Price,
		Satuan:    p.Satuan,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func toCoreList(resp []Product) []product.Core {
	p := []product.Core{}
	for key := range resp {
		p = append(p, resp[key].toCore())
	}
	return p
}

func fromCore(core product.Core) Product {
	return Product{}
}
