package factory

import (
	"A-Golang-MiniProject/config"
	_distributor_bussiness "A-Golang-MiniProject/features/distributor/bussiness"
	_distributor_data "A-Golang-MiniProject/features/distributor/data"
	_distributor_presentation "A-Golang-MiniProject/features/distributor/presentation"
	_product_bussiness "A-Golang-MiniProject/features/product/bussiness"
	_product_data "A-Golang-MiniProject/features/product/data"
	_product_presentation "A-Golang-MiniProject/features/product/presentation"
)

type Presenter struct {
	ProductPresentation *_product_presentation.ProductHandler
}

func Init() Presenter {
	productData := _product_data.NewProductRepository(config.DB)
	productBussiness := _product_bussiness.NewProductBussiness(productData)
	productPresentation := _product_presentation.NewProductHandler(productBussiness)

	distributorData := _distributor_data.NewDistribRepository(config.DB)
	distributorBussiness := _distributor_bussiness.NewDistributorBussiness(distributorData)
	distributorPresentation := _distributor_presentation.NewDistribHandler(distributorBussiness)

	return Presenter{
		ProductPresentation:     productPresentation,
		DistributorPresentation: distributorPresentation,
	}
}
