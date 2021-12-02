package presentation

import (
	"A-Golang-MiniProject/features/spoonacular"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApiResponse struct {
	HealtyFood interface{} `json:"healtyFood"`
}

func FromDomain(core spoonacular.Core) ApiResponse {
	return ApiResponse{
		HealtyFood: core.HealtyFood,
	}
}

func FromListDomain(domain []spoonacular.Core) []ApiResponse {
	var response []ApiResponse
	for _, value := range domain {
		response = append(response, FromDomain(value))
	}
	return response
}

type BaseResponse struct {
	Meta struct {
		Message string `json:"message"`
	} `json:"meta"`
	HealtyFood interface{} `json:"HealtyFood"`
}

func NewSuccesResponse(c echo.Context, param interface{}) error {
	response := BaseResponse{}
	response.Meta.Message = "You lucky today this apps running"
	response.HealtyFood = param
	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Message = "Yesterday this app running successfully, but I dont know why"
	return c.JSON(status, response)
}
