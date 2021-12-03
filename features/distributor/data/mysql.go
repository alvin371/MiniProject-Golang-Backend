package data

import (
	"A-Golang-MiniProject/features/distributor"

	"gorm.io/gorm"
)

type mysqlDistributorRepository struct {
	Conn *gorm.DB
}

func NewDistribRepository(conn *gorm.DB) distributor.Data {
	return &mysqlDistributorRepository{
		Conn: conn,
	}
}

func (db *mysqlDistributorRepository) GetAllData(name string) (resp []distributor.Core) {
	var record []Distributor
	if err := db.Conn.Preload("Distributor").Find(&record).Error; err != nil {
		return []distributor.Core{}
	}
	return toCoreList(record)
}

func (db *mysqlDistributorRepository) InsertData(data distributor.Core) (resp distributor.Core, err error) {
	record := fromCore(data)

	if err := db.Conn.Create(&record).Error; err != nil {
		return distributor.Core{}, err
	}
	return data, nil
}
