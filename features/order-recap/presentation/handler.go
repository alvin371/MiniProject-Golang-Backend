package presentation

import (
	"A-Golang-MiniProject/features/order-recap"
	"A-Golang-MiniProject/features/order-recap/presentation/request"
	"A-Golang-MiniProject/features/order-recap/presentation/response"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	ob order.Bussiness
}

func NewHandlerOrder(ob order.Bussiness) *OrderHandler {
	return &OrderHandler{ob}
}

func (oh *OrderHandler) Create(c echo.Context) error {
	req := request.Order{}
	fmt.Print(req)
	if err := c.Bind(&req); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := oh.ob.Create(req.ToDomain())

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func (oh *OrderHandler) GetAll(c echo.Context) error {
	result, err := oh.ob.GetAll()

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccesResponse(c, response.FromProductListDomain(result))
}
