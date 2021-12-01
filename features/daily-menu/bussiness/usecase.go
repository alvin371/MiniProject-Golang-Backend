package bussiness

import (
	dailymenu "A-Golang-MiniProject/features/daily-menu"
)

type DailyMenuUseCase struct {
	DailyMenu dailymenu.Data
}

func NewDailyMenuBussiness(dayData dailymenu.Data) dailymenu.Bussiness {
	return &DailyMenuUseCase{
		DailyMenu: dayData,
	}
}

func (dy *DailyMenuUseCase) GetAllData(search string) (resp []dailymenu.DailyMenu) {
	resp = dy.DailyMenu.SelectData(search)
	return
}

func (dy *DailyMenuUseCase) CreateData(data dailymenu.DailyMenu) (resp dailymenu.DailyMenu, err error) {
	resp, err = dy.DailyMenu.InsertData(data)
	if err != nil {
		return dailymenu.DailyMenu{}, err
	}
	return dailymenu.DailyMenu{}, err
}
