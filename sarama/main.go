package main

import (
	"example-sarama/app/interface/container"
	"example-sarama/app/interface/server"
)


func main() {
	server.StartServer(container.SetupContainer())
}