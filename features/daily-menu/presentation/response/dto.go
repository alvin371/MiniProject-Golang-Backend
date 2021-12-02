package response

import (
	dailymenu "A-Golang-MiniProject/features/daily-menu"
	"time"
)

type DailyMenu struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	Desc      string    `json:"Desc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromCore(core dailymenu.Core) DailyMenu {
	return DailyMenu{
		ID:        core.ID,
		Name:      core.Name,
		Price:     core.Price,
		Desc:      core.Desc,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
	}
}

func FromCoreSlice(Core []dailymenu.Core) []DailyMenu {
	var artArray []DailyMenu
	for key := range Core {
		artArray = append(artArray, FromCore(Core[key]))
	}
	return artArray
}
