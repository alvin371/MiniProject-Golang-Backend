package main

import (
	"A-Golang-MiniProject/config"
	"A-Golang-MiniProject/migrate"
	"A-Golang-MiniProject/routes"
)

func main() {
	// DB
	config.InitDB()

	migrate.AutoMigrate()

	//initRoutes
	e := routes.New()

	//starting the server
	e.Start(":8000")
}
