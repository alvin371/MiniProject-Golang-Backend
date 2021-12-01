package dailymenu

import (
	"time"
)

type DailyMenu struct {
	ID        int
	Name      string
	Price     float32
	Obj       []Obj
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Obj struct {
	ID   int
	Name string
}

type Bussiness interface {
	GetAllData(search string) (resp []DailyMenu)
	CreateData(data DailyMenu) (resp DailyMenu, err error)
}

type Data interface {
	InsertData(data DailyMenu) (resp DailyMenu, err error)
	SelectData(name string) (resp []DailyMenu)
}
