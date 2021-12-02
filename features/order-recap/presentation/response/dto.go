package response

import (
	"A-Golang-MiniProject/features/order-recap"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type CreateOrderResponse struct {
	Message     string    `json:"message"`
	ID          int       `json:"id"`
	Total_Price int       `json:"total_price"`
	Worker      string    `json:"worker_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type GetAllOrderResponse struct {
	ID          int       `json:"id"`
	Total_Price int       `json:"total_price"`
	Worker      string    `json:"worker_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BasicResponse struct {
	Meta struct {
		Message string `json:"message"`
	}
	Data interface{} `json:"data"`
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BasicResponse{}
	response.Meta.Message = "You not lucky today, our apps sucks"
	return c.JSON(status, response)
}

func NewSuccesResponse(c echo.Context, param interface{}) error {
	response := BasicResponse{}
	response.Meta.Message = "You lucky today our apps running"
	return c.JSON(http.StatusOK, response)
}

func FromDomainGetAll(ok order.Core) GetAllOrderResponse {
	return GetAllOrderResponse{
		ID:          ok.ID,
		Total_Price: ok.Total_Price,
		Worker:      ok.Worker,
		CreatedAt:   ok.CreatedAt,
		UpdatedAt:   ok.UpdatedAt,
	}
}

func FromDomainCreate(ok order.Core) CreateOrderResponse {
	return CreateOrderResponse{
		Message:     "You lucky today our apps running",
		ID:          ok.ID,
		Total_Price: ok.Total_Price,
		Worker:      ok.Worker,
		CreatedAt:   ok.CreatedAt,
		UpdatedAt:   ok.UpdatedAt,
	}
}

func FromProductListDomain(domain []order.Core) []GetAllOrderResponse {
	var response []GetAllOrderResponse
	for _, value := range domain {
		response = append(response, FromDomainGetAll(value))
	}
	return response
}
