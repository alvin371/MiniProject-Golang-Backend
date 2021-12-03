package routes

import (
	"github.com/labstack/echo/v4"

	"A-Golang-MiniProject/factory"
)

func New() *echo.Echo {

	presenter := factory.Init()

	//initiate
	e := echo.New()
	// V1 := e.Group("v1/api/")
	e.GET("/product", presenter.ProductPresentation.GetAllData)
	e.POST("/product", presenter.ProductPresentation.AddProduct)

	e.GET("/order-recap", presenter.OrderPresentation.GetAll)
	e.POST("/order-recap", presenter.OrderPresentation.Create)

	return e
}
