package presentation

import (
	"A-Golang-MiniProject/features/product"
	"net/http"

	presentation_response "A-Golang-MiniProject/features/product/presentation/response"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productBussiness product.Bussiness
}

func (ph *ProductHandler) GetAllData(c echo.Context) error {
	result := ph.productBussiness.GetAllData("")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "All Data Printed",
		"data":    presentation_response.FromCoreSlice(result),
	})
}
