package main

import (
	"tugas_2/database"
	"tugas_2/routers"
)

func main() {

	port := ":7000"

	database.GormInit()

	r := routers.InitRouter()

	// Run the server
	r.Run(port)
}
