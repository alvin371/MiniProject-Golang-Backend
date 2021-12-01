package bussiness

import (
	"A-Golang-MiniProject/features/product"
)

type productUseCase struct {
	productData product.Data
}

func NewProductBussiness(prodData product.Data) product.Bussiness {
	return &productUseCase{
		productData: prodData,
	}
}

func (pu *productUseCase) GetAllData(search string) (resp []product.Core) {
	resp = pu.productData.SelectData(search)
	return

}

func (pu *productUseCase) CreateData(data product.Core) (resp product.Core, err error) {
	resp, err = pu.productData.InsertData(data)
	if err != nil {
		return product.Core{}, err
	}
	return product.Core{}, err
}
