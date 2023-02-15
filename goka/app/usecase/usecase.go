package usecase

import "github.com/lovoo/goka"

type Usecase interface {
	Notif() (res func(ctx goka.Context, msg interface{}), err error)
}
