package data

import (
	dm "A-Golang-MiniProject/features/daily-menu"

	"gorm.io/gorm"
)

type mysqlDailyMenuRepository struct {
	Conn *gorm.DB
}

func NewmysqlDailyMenuRepository(conn *gorm.DB) dm.Data {
	return &mysqlDailyMenuRepository{
		Conn: conn,
	}
}

func (rep *mysqlDailyMenuRepository) CreateData(domain dm.Core) (dm.Core, error) {
	obj := fromDomain(domain)

	result := rep.Conn.Create(&obj)
	if result.Error != nil {
		return dm.Core{}, result.Error
	}
	return obj.toDomain(), nil
}

func (rep *mysqlDailyMenuRepository) GetAllData(name string) (resp []dm.Core) {
	var record []DailyMenu
	if err := rep.Conn.Preload("DailyMenu").Find(&record).Error; err != nil {
		return []dm.Core{}
	}
	return toCoreList(record)
}
