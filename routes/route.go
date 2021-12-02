package routes

import (
	"github.com/labstack/echo/v4"

	"A-Golang-MiniProject/factory"
)

func New() *echo.Echo {

	presenter := factory.Init()

	//initiate
	e := echo.New()
	e.GET("/product", presenter.ProductPresentation.GetAllData)
	e.POST("/product", presenter.ProductPresentation.AddProduct)
	e.GET("/distributor", presenter.DistributorPresentation.GetAllData)
	e.POST("/distributor", presenter.DistributorPresentation.CreateData)

	return e
}
