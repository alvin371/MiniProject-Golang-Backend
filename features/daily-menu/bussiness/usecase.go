package bussiness

import dm "A-Golang-MiniProject/features/daily-menu"

type DailyMenuUseCase struct {
	DailyMenu dm.Data
}

func NewDailyMenuBussiness(dayData dm.Data) dm.Bussiness {
	return &DailyMenuUseCase{
		DailyMenu: dayData,
	}
}

func (dy *DailyMenuUseCase) GetAllData(search string) (resp []dm.Core) {
	resp = dy.DailyMenu.GetAllData(search)
	return
}

func (dy *DailyMenuUseCase) CreateData(data dm.Core) (resp dm.Core, err error) {
	resp, err = dy.DailyMenu.CreateData(data)
	if err != nil {
		return dm.Core{}, err
	}
	return dm.Core{}, err
}
