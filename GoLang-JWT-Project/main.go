package main

import (
	"gin-mongo-api/db"
	"gin-mongo-api/routes"
)

func main() {
	db.ConnectToDB()
	db.SyncDatabase()
	routes.Routers()
}
