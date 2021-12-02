package dailymenu

import (
	"time"
)

type DailyMenu struct {
	ID        int
	Name      string
	Price     float32
	Desc      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Bussiness interface {
	GetAllData(search string) (resp []DailyMenu)
	CreateData(data *DailyMenu) (resp DailyMenu, err error)
}

type Data interface {
	SelectData(name string) (resp []DailyMenu)
	InsertData(data *DailyMenu) (resp DailyMenu, err error)
}
