package dm

import (
	"time"
)

type Core struct {
	ID        int
	Name      string
	Price     float32
	Desc      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Bussiness interface {
	GetAllData(search string) (resp []Core)
	CreateData(data Core) (resp Core, err error)
}

type Data interface {
	GetAllData(name string) (resp []Core)
	CreateData(data Core) (resp Core, err error)
}
