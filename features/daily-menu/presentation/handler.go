package presentation

import (
	dailymenu "A-Golang-MiniProject/features/daily-menu"
	presentation_request "A-Golang-MiniProject/features/daily-menu/presentation/request"
	presentation_response "A-Golang-MiniProject/features/daily-menu/presentation/response"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
)

type DailyMenuHandler struct {
	dmBussiness dailymenu.Bussiness
}

func NewDMHandler(dm dailymenu.Bussiness) *DailyMenuHandler {
	return &DailyMenuHandler{
		dmBussiness: dm,
	}
}

// get data

func (dm *DailyMenuHandler) GetAllData(c echo.Context) error {
	result := dm.dmBussiness.GetAllData("")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "You are lucky today, All Data Printed while program have no bugs",
		"data":    presentation_response.FromCoreSlice(result),
	})
}

// insert data

func (dm *DailyMenuHandler) AddDailyMenu(c echo.Context) error {
	var NewDM presentation_request.DailyMenu
	c.Bind(&NewDM)
	resp, err := dm.dmBussiness.CreateData(presentation_request.ToCore(NewDM))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Data Created",
		"data":    presentation_response.FromCore(resp),
	})
}
