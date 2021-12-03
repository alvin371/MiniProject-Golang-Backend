package presentation

import (
	"A-Golang-MiniProject/features/distributor"
	presentation_request "A-Golang-MiniProject/features/distributor/presentation/request"
	presentation_response "A-Golang-MiniProject/features/distributor/presentation/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DistribHandler struct {
	distribBussiness distributor.Bussiness
}

func NewDistribHandler(db distributor.Bussiness) *DistribHandler {
	return &DistribHandler{
		distribBussiness: db,
	}
}

func (db *DistribHandler) GetAllData(c echo.Context) error {
	result := db.distribBussiness.GetAllData("")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "This is your lucky day, coz this app have no bug",
		"data":    presentation_response.FromCoreSlice(result),
	})
}

func (db *DistribHandler) CreateData(c echo.Context) error {
	var newDB presentation_request.Distributor
	c.Bind(&newDB)
	resp, err := db.distribBussiness.CreateData(presentation_request.ToCore(newDB))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Alhamdulillah datanya Masuk",
		"data":    presentation_response.FromCore(resp),
	})
}
