package main

import (
	"chatserver/database"
	"chatserver/server"
	"chatserver/user"
)

func main() {
	user.Service.Start()
	database.Service.Start()
	server.Service.Start()
}
