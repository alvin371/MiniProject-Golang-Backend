package distributor

import (
	"time"
)

type Core struct {
	ID        int
	Name      string
	Telp      float32
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Bussiness interface {
	GetAllData(search string) (resp []Core)
	CreateData(data Core) (resp Core, err error)
}

type Data interface {
	GetAllData(name string) (resp []Core)
	InsertData(data Core) (resp Core, err error)
}
