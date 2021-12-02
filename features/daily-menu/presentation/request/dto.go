package request

import dailymenu "A-Golang-MiniProject/features/daily-menu"

type DailyMenu struct {
	Name  string
	Price float32
	Desc  string
}

func (req *DailyMenu) ToCore() *dailymenu.DailyMenu {
	return &dailymenu.DailyMenu{
		Name:  req.Name,
		Price: req.Price,
		Desc:  req.Desc,
	}
}
