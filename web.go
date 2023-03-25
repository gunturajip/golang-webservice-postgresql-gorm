package main

import (
	"golang-webservice/database"
	"golang-webservice/routers"
)

var PORT = "127.0.0.1:8080"

func main() {
	database.StartDB()
	routers.StartServer().Run(PORT)
}
