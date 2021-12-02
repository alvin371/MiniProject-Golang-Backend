package factory

import (
	"A-Golang-MiniProject/config"
	"A-Golang-MiniProject/features/order-recap"
	orderDB "A-Golang-MiniProject/features/order-recap/data"
	_product_bussiness "A-Golang-MiniProject/features/product/bussiness"
	_product_data "A-Golang-MiniProject/features/product/data"
	_product_presentation "A-Golang-MiniProject/features/product/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	ProductPresentation *_product_presentation.ProductHandler
}

func Init() Presenter {
	productData := _product_data.NewProductRepository(config.DB)
	productBussiness := _product_bussiness.NewProductBussiness(productData)
	productPresentation := _product_presentation.NewProductHandler(productBussiness)

	return Presenter{
		ProductPresentation: productPresentation,
	}
}

func NewOrderRepository(conn *gorm.DB) order.Data {
	return orderDB.NewMySqlOrderRepository(conn)
}
