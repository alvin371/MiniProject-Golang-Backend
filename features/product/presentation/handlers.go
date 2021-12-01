package presentation

import (
	"A-Golang-MiniProject/features/product"
	"net/http"

	presentation_request "A-Golang-MiniProject/features/product/presentation/request"
	presentation_response "A-Golang-MiniProject/features/product/presentation/response"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productBussiness product.Bussiness
}

func NewProductHandler(pb product.Bussiness) *ProductHandler {
	return &ProductHandler{
		productBussiness: pb,
	}
}

// get data

func (ph *ProductHandler) GetAllData(c echo.Context) error {
	result := ph.productBussiness.GetAllData("")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "All Data Printed",
		"data":    presentation_response.FromCoreSlice(result),
	})
}

// insert data

func (ph *ProductHandler) AddProduct(c echo.Context) error {
	var newProduct presentation_request.Product
	c.Bind(&newProduct)
	resp, err := ph.productBussiness.CreateData(presentation_request.ToCore(newProduct))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Data Created",
		"data":    presentation_response.FromCore(resp),
	})
}
