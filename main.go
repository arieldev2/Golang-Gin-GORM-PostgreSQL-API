package main

import (
	"gin_test/migrations"
	"gin_test/routes"
)

func main() {

	migrations.InitialMigration()

	routes.Routes()

}
