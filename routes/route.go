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

	return e
}
