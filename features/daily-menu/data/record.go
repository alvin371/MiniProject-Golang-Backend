package data

import (
	dm "A-Golang-MiniProject/features/daily-menu"

	"gorm.io/gorm"
)

type DailyMenu struct {
	gorm.Model
	Name  string `gorm:"type:varchar(50)"`
	Price float32
	Desc  string
}

func (d *DailyMenu) toDomain() dm.Core {
	return dm.Core{
		ID:        int(d.ID),
		Name:      d.Name,
		Price:     d.Price,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func toCoreList(resp []DailyMenu) []dm.Core {
	dm := []dm.Core{}
	for key := range resp {
		dm = append(dm, resp[key].toDomain())
	}
	return dm
}

func fromDomain(core dm.Core) DailyMenu {
	return DailyMenu{}
}
