package order

import "time"

type Core struct {
	ID          int
	Total_Price int
	Worker      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Bussiness interface {
	GetAll() ([]Core, error)
	Create(core *Core) (Core, error)
}

type Data interface {
	GetAll() ([]Core, error)
	Create(core *Core) (Core, error)
}
