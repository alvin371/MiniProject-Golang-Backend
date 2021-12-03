package bussiness

import "A-Golang-MiniProject/features/order-recap"

type serviceOrder struct {
	orderRepo order.Data
}

func NewServiceOrder(repoOrder order.Data) order.Bussiness {
	return &serviceOrder{orderRepo: repoOrder}
}

func (so *serviceOrder) GetAll() ([]order.Core, error) {
	result, err := so.orderRepo.GetAll()

	if err != nil {
		return []order.Core{}, err
	}
	return result, nil
}

func (so *serviceOrder) Create(core order.Core) error {
	err := so.orderRepo.Create(core)
	if err != nil {
		return err
	}
	return nil
}
