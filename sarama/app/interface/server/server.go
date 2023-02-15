package server

import (
	"example-sarama/app/interface/container"
)

func StartServer(container container.Container) {
	// container.Event.Run()
	container.Event.Notif()
}
