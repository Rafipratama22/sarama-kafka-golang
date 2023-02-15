package usecase

import "example-sarama/app/usecase/request"

type Usecase interface {
	Notif(in *request.NotifRequest) (res string, err error)
}
