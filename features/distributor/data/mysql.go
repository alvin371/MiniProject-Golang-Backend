package data

import (
	"A-Golang-MiniProject/features/distributor"

	"gorm.io/gorm"
)

type mysqlDistributorRepository struct {
	DB *gorm.DB
}

func NewDistribRepository(conn *gorm.DB) distributor.Data {
	return &mysqlDistributorRepository{
		DB: conn,
	}
}

func (db *mysqlDistributorRepository) GetAllData(name string) (resp []distributor.Core) {
	var record []Distributor
	if err := db.DB.Preload("Distributor").Find(&record).Error; err != nil {
		return []distributor.Core{}
	}
	return toCoreList(record)
}

func (db *mysqlDistributorRepository) CreateData(data distributor.Core) (resp distributor.Core, err error) {
	record := fromCore(data)

	if err := db.DB.Create(&record).Error; err != nil {
		return distributor.Core{}, err
	}
	return data, nil
}
