package data

import (
	"A-Golang-MiniProject/features/distributor"

	"gorm.io/gorm"
)

type Distributor struct {
	gorm.Model
	Name    string `gorm:"type:varchar(25)"`
	Telp    float32
	Address string
}

func (d *Distributor) toCore() distributor.Core {
	return distributor.Core{
		ID:        int(d.ID),
		Name:      d.Name,
		Telp:      d.Telp,
		Address:   d.Address,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func toCoreList(resp []Distributor) []distributor.Core {
	d := []distributor.Core{}
	for key := range resp {
		d = append(d, resp[key].toCore())
	}
	return d
}

func fromCore(core distributor.Core) Distributor {
	return Distributor{}
}
