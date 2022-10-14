package main

import (
	"gin-mongo-api/db"
	"gin-mongo-api/routes"
)

func main() {
	go db.ConnectToDB()
	routes.Routers()
}
