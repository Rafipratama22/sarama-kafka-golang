package transport

import (
	"example-goka/app/interface/container"
)

type Tp struct {
	Transport *transport
}

func SetupTransport(container *container.Container) (t *Tp) {
	tp := NewTransport(container)
	return &Tp{
		Transport: tp,
	}
}
