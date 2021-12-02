package request

import (
	dm "A-Golang-MiniProject/features/daily-menu"
)

type DailyMenu struct {
	Name  string
	Price float32
	Desc  string
}

func (req DailyMenu) ToDomain() dm.Core {
	return dm.Core{
		Name:  req.Name,
		Price: req.Price,
		Desc:  req.Desc,
	}
}

func ToCore(req DailyMenu) dm.Core {
	return dm.Core{
		Name:  req.Name,
		Price: req.Price,
		Desc:  req.Desc,
	}
}
