package presentation

import (
	"A-Golang-MiniProject/features/spoonacular"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApiHandler struct {
	apiRepo spoonacular.Data
}

func NewApiHandler(apiRepo spoonacular.Data) *ApiHandler {
	return &ApiHandler{
		apiRepo: apiRepo,
	}
}

func (ah ApiHandler) GetByNutrients(c echo.Context) error {
	random := true
	var minCarbs int32 = 50
	var number int16 = 1
	data, error := ah.apiRepo.GetByNutrients(random, minCarbs, number)

	if error != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, error)
	}
	return NewSuccesResponse(c, FromDomain(data))
}
