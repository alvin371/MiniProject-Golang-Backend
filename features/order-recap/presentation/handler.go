package presentation

import (
	"A-Golang-MiniProject/features/order-recap"
	"A-Golang-MiniProject/features/order-recap/presentation/request"
	"A-Golang-MiniProject/features/order-recap/presentation/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderBussiness order.Bussiness
}

func NewHandlerOrder(serv order.Bussiness) *OrderHandler {
	return &OrderHandler{
		orderBussiness: serv,
	}
}

func (oh *OrderHandler) Create(c echo.Context) error {
	req := request.Order{}

	if err := c.Bind(&req); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := oh.orderBussiness.Create(req.ToDomain())

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccesResponse(c, response.FromDomainCreate(result))
}

func (oh *OrderHandler) GetAll(c echo.Context) error {
	result, err := oh.orderBussiness.GetAll()

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccesResponse(c, response.FromProductListDomain(result))
}
