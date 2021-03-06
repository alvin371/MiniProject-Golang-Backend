package product

import "time"

type Core struct {
	ID          int
	Name        string
	Price       float32
	Satuan      int
	Distributor []Distributor
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Distributor struct {
	Name    string
	Telp    int
	Address string
}

type Bussiness interface {
	GetAllData(search string) (resp []Core)
	CreateData(data Core) (resp Core, err error)
}
type Data interface {
	SelectData(name string) (resp []Core)
	InsertData(data Core) (resp Core, err error)
}
