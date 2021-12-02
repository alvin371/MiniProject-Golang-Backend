package data

import (
	dailymenu "A-Golang-MiniProject/features/daily-menu"

	"gorm.io/gorm"
)

type mysqlDailyMenuRepository struct {
	Conn *gorm.DB
}

func NewmysqlDailyMenuRepository(conn *gorm.DB) dailymenu.Data {
	return &mysqlDailyMenuRepository{
		Conn: conn,
	}
}

func (rep mysqlDailyMenuRepository) CreateData(domain *dailymenu.DailyMenu) (dailymenu.DailyMenu, error) {
	obj := fromDomain(domain)

	result := rep.Conn.Create(&obj)
	if result.Error != nil {
		return dailymenu.DailyMenu{}, result.Error
	}
	return toDomain(obj), nil
}

func (rep *mysqlDailyMenuRepository) SelectData(name string) (resp []dailymenu.DailyMenu) {
	var record []dailymenu.DailyMenu
	if err := rep.Conn.Preload("DailyMenu").Find(&record).Error; err != nil {
		return []dailymenu.DailyMenu{}
	}
	return toCoreList(record)
}
