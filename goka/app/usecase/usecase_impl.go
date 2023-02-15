package usecase

import (
	"log"

	"example-goka/app/model/proto/notif"

	"github.com/lovoo/goka"
	"google.golang.org/protobuf/proto"
)

type usecase struct {
}

func NewUsecase() Usecase {
	return &usecase{}
}

func (u *usecase) Notif() (res func(ctx goka.Context, msg interface{}), err error) {
	return func(ctx goka.Context, msg interface{}) {
		check := &notif.NotifRequest{}
		_ = proto.Unmarshal(msg.([]byte), check)
		log.Println(check)
	}, nil

}
