package bussiness

import (
	"A-Golang-MiniProject/features/distributor"
)

type distributorUSeCase struct {
	distribData distributor.Data
}

func NewDistributorBussiness(DbData distributor.Data) distributor.Bussiness {
	return &distributorUSeCase{
		distribData: DbData,
	}
}

func (db *distributorUSeCase) GetAllData(search string) (resp []distributor.Core) {
	resp = db.distribData.GetAllData(search)
	return
}

func (db *distributorUSeCase) CreateData(data distributor.Core) (resp distributor.Core, err error) {
	resp, err = db.distribData.InsertData(data)
	if err != nil {
		return distributor.Core{}, err
	}
	return distributor.Core{}, err
}
