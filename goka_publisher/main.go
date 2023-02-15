package main

import (
	"example-goka/app/interface/container"
	"example-goka/app/interface/server"
)

func main() {
	server.StartServer(container.SetupContainer())
}
