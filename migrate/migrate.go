package migrate

import (
	"A-Golang-MiniProject/config"
	m_daily "A-Golang-MiniProject/features/daily-menu/data"
	m_distributor "A-Golang-MiniProject/features/distributor/data"
	m_product "A-Golang-MiniProject/features/product/data"
)

func AutoMigrate() {
	config.DB.AutoMigrate(
		&m_product.Product{},
		&m_product.Distributor{},
		&m_distributor.Distributor{},
		&m_daily.DailyMenu{},
	)
}
