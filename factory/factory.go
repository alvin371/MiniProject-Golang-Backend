package factory

import (
	"A-Golang-MiniProject/config"
	orderB "A-Golang-MiniProject/features/order-recap/bussiness"
	orderDB "A-Golang-MiniProject/features/order-recap/data"
	orderP "A-Golang-MiniProject/features/order-recap/presentation"
	_product_bussiness "A-Golang-MiniProject/features/product/bussiness"
	_product_data "A-Golang-MiniProject/features/product/data"
	_product_presentation "A-Golang-MiniProject/features/product/presentation"
)

type Presenter struct {
	ProductPresentation *_product_presentation.ProductHandler
	OrderPresentation   *orderP.OrderHandler
}

func Init() Presenter {
	productData := _product_data.NewProductRepository(config.DB)
	productBussiness := _product_bussiness.NewProductBussiness(productData)
	productPresentation := _product_presentation.NewProductHandler(productBussiness)

	orderData := orderDB.NewMySqlOrderRepository(config.DB)
	orderBussiness := orderB.NewServiceOrder(orderData)
	orderPresentation := orderP.NewHandlerOrder(orderBussiness)
	return Presenter{
		ProductPresentation: productPresentation,
		OrderPresentation:   orderPresentation,
	}
}
