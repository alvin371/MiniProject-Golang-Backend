package factory

import (
	"A-Golang-MiniProject/config"
	_daily_bussiness "A-Golang-MiniProject/features/daily-menu/bussiness"
	_daily_data "A-Golang-MiniProject/features/daily-menu/data"
	_daily_presentation "A-Golang-MiniProject/features/daily-menu/presentation"
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
	DailyPresentation   *_daily_presentation.DailyMenuHandler
}

func Init() Presenter {
	productData := _product_data.NewProductRepository(config.DB)
	productBussiness := _product_bussiness.NewProductBussiness(productData)
	productPresentation := _product_presentation.NewProductHandler(productBussiness)

	orderData := orderDB.NewMySqlOrderRepository(config.DB)
	orderBussiness := orderB.NewServiceOrder(orderData)
	orderPresentation := orderP.NewHandlerOrder(orderBussiness)
	// distributorData := _distributor_data.NewDistribRepository(config.DB)
	// distributorBussiness := _distributor_bussiness.NewDistributorBussiness(distributorData)
	// distributorPresentation := _distributor_presentation.NewDistribHandler(distributorBussiness)

	dailyData := _daily_data.NewmysqlDailyMenuRepository(config.DB)
	dailyBussiness := _daily_bussiness.NewDailyMenuBussiness(dailyData)
	dailyPresentation := _daily_presentation.NewDMHandler(dailyBussiness)

	return Presenter{
		ProductPresentation: productPresentation,
		OrderPresentation:   orderPresentation,
		DailyPresentation:   dailyPresentation,
		// DistributorPresentation: distributorPresentation,
	}
}
