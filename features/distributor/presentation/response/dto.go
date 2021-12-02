package response

import (
	"A-Golang-MiniProject/features/distributor"
	"time"
)

type Distributor struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Telp      float32   `json:"telp"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromCore(core distributor.Core) Distributor {
	return Distributor{
		ID:        core.ID,
		Name:      core.Name,
		Telp:      core.Telp,
		Address:   core.Address,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
	}
}

func FromCoreSlice(core []distributor.Core) []Distributor {
	var artArray []Distributor
	for key := range core {
		artArray = append(artArray, FromCore(core[key]))
	}
	return artArray
}
