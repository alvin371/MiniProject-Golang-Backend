package routes

import (
	"A-Golang-MiniProject/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	our_middleware "A-Golang-MiniProject/middleware"
)

func New() *echo.Echo {

	presenter := factory.Init()

	//initiate
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	our_middleware.LogMiddleware(e)

	e.GET("/product", presenter.ProductPresentation.GetAllData)
	e.POST("/product", presenter.ProductPresentation.AddProduct)
	// e.GET("/distributor", presenter.DistributorPresentation.GetAllData)
	// e.POST("/distributor", presenter.DistributorPresentation.CreateData)

	e.GET("/daily-menu", presenter.ProductPresentation.GetAllData)
	e.POST("/daily-menu", presenter.ProductPresentation.AddProduct)

	return e
}
