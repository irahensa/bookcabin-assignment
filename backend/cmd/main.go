package main

import (
	"github.com/irahensa/bookcabin-assignment/backend/config"
	"github.com/irahensa/bookcabin-assignment/backend/server"
)

func main() {
	config.InitConfig()

	server.Serve()
}
