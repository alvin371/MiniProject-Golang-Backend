package routes

import (
	"github.com/labstack/echo/v4"

	"A-Golang-MiniProject/factory"
	order "A-Golang-MiniProject/features/order-recap/presentation"
)

type RouteList struct {
	OrderRouter order.OrderHandler
}

func (rt *RouteList) New() *echo.Echo {

	presenter := factory.Init()

	//initiate
	e := echo.New()

	V1 := e.Group("v1/api/")
	V1.GET("/product", presenter.ProductPresentation.GetAllData)
	V1.POST("/product", presenter.ProductPresentation.AddProduct)

	V1.GET("/order-recap", rt.OrderRouter.GetAll)
	V1.POST("/order-recap", rt.OrderRouter.Create)

	return e
}
