package usecase

import "example-goka/app/usecase/request"

type Usecase interface {
	Notif(in *request.NotifRequest) (res string, err error)
}
