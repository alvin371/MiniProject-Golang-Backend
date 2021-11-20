package bussiness

import (
	"A-Golang-MiniProject/features/product"
)

type productUseCase struct {
	productData product.Data
}

func (pu *productUseCase) GetAllData(search string) (resp []product.Core) {
	resp = pu.productData.SelectData(search)
	return

}
