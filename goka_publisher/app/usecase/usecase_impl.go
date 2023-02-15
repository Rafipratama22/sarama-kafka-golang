package usecase

import (
	"log"

	"example-goka/app/model/proto/notif"
	"example-goka/app/repository"
	"example-goka/app/usecase/request"

	"github.com/lovoo/goka"
	"google.golang.org/protobuf/proto"
)

type usecase struct {
	producer *goka.Emitter
	repo     repository.Repository
}

func NewUsecase(producer *goka.Emitter, repository repository.Repository) Usecase {
	return &usecase{
		producer: producer,
		repo:     repository,
	}
}

func (u *usecase) Notif(in *request.NotifRequest) (res string, err error) {
	log.Println("masukk")

	topic := "test_topic"

	notifStruct := &notif.NotifRequest{
		Title:            in.Data.Title,
		Body:             in.Data.Body,
		To:               in.To,
		Image:            "check",
		Channel:          in.Data.Channel,
		NotificationId:   in.NotificationId,
		NotificationType: in.NotificationType,
	}
	rest, err := proto.Marshal(notifStruct)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rest)

	SendMessage(topic, rest, u.producer)
	return "berhasil", nil
}

func SendMessage(topic string, message []byte, producer *goka.Emitter) {
	err := producer.EmitSync("some_key", message)
	if err != nil {
		log.Fatal(err)
	}
}
