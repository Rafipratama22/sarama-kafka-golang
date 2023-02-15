package server

import (
	"example-goka/app/interface/container"
)

func StartServer(container container.Container) {
	// container.Event.Run()
	container.Event.Notif()
}
