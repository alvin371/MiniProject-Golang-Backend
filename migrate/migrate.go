package migrate

import (
	"A-Golang-MiniProject/config"
	m_product "A-Golang-MiniProject/features/product/data"
)

func AutoMigrate() {
	config.DB.AutoMigrate(
		&m_product.Product{},
		&m_product.Distributor{},
	)
}
