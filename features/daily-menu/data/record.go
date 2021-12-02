package data

import (
	dailymenu "A-Golang-MiniProject/features/daily-menu"

	"gorm.io/gorm"
)

type DailyMenu struct {
	gorm.Model
	Name  string `gorm:"type:varchar(50)"`
	Price float32
	Desc  string
}

func (dm *DailyMenu) toCore() dailymenu.DailyMenu {
	return dailymenu.DailyMenu{
		ID:        int(dm.ID),
		Name:      dm.Name,
		Price:     dm.Price,
		CreatedAt: dm.CreatedAt,
		UpdatedAt: dm.UpdatedAt,
	}
}

func toCoreList(resp []DailyMenu) []dailymenu.DailyMenu {
	dm := []dailymenu.DailyMenu{}
	for key := range resp {
		dm = append(dm, resp[key].toCore())
	}
	return dm
}

func fromDomain(core dailymenu.DailyMenu) DailyMenu {
	return DailyMenu{}
}
