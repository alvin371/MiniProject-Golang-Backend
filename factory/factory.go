package factory

import (
	"A-Golang-MiniProject/config"
	_product_bussiness "A-Golang-MiniProject/features/products/bussiness"
	_product_data "A-Golang-MiniProject/features/products/data"
	_product_presentation "A-Golang-MiniProject/features/products/presentation"
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
