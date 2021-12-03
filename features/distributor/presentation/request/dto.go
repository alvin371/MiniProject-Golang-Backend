package request

import "A-Golang-MiniProject/features/distributor"

type Distributor struct {
	Name    string  `json:"name"`
	Telp    float32 `json:"telp"`
	Address string  `json:"address"`
}

func ToCore(req Distributor) distributor.Core {
	return distributor.Core{
		Name:    req.Name,
		Telp:    req.Telp,
		Address: req.Address,
	}
}
