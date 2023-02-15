package usecase

import (
	"log"

	"example-sarama/app/model/proto/notif"
	"example-sarama/app/repository"
	"example-sarama/app/usecase/request"

	"github.com/Shopify/sarama"
	"google.golang.org/protobuf/proto"
)

type usecase struct {
	producer sarama.SyncProducer
	repo repository.Repository
}

func NewUsecase(producer sarama.SyncProducer, repository repository.Repository) Usecase {
	return &usecase{
		producer: producer,
		repo: repository,
	}
}


func (u *usecase) Notif(in *request.NotifRequest) (res string, err error) {
	log.Println("masukk")

	topic := "test_topic"

	notifStruct := &notif.NotifRequest{
		Title: in.Data.Title,
		Body: in.Data.Body,
		To: in.To,
		Image: "check",
		Channel: in.Data.Channel,
		NotificationId: in.NotificationId,
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


func SendMessage(topic string, message []byte, producer sarama.SyncProducer) {
	kafka := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err :=producer.SendMessage(kafka)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(partition, offset)
}