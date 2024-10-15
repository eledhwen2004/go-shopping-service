package main

import (
	"shopping-clone/api"
	"shopping-clone/migration"
	"shopping-clone/postgre"
)

func main() {
	postgre.ConnectDatabase()
	migration.MigrateDB()
	api.ListenPort()
}
