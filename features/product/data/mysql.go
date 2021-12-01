package data

import (
	"A-Golang-MiniProject/features/product"

	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	Conn *gorm.DB
}

func NewProductRepository(conn *gorm.DB) product.Data {
	return &mysqlProductRepository{
		Conn: conn,
	}
}

func (ar *mysqlProductRepository) SelectData(title string) (resp []product.Core) {
	var record []Product
	if err := ar.Conn.Preload("Distributor").Find(&record).Error; err != nil {
		return []product.Core{}
	}
	return toCoreList(record)
}

func (pr *mysqlProductRepository) InsertData(data product.Core) (resp product.Core, err error) {
	record := fromCore(data)

	if err := pr.Conn.Create(&record).Error; err != nil {
		return product.Core{}, err
	}
	return data, nil
}
