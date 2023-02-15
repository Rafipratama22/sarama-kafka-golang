package usecase

import (
	"log"

	"example-sarama/app/model/proto/notif"

	"github.com/Shopify/sarama"
	"google.golang.org/protobuf/proto"
)

type usecase struct {
	consumer sarama.Consumer
}

func NewUsecase(consumer sarama.Consumer) Usecase {
	return &usecase{
		consumer: consumer,
	}
}


func (u *usecase) Notif() {
	chanMessage := make(chan *sarama.ConsumerMessage, 256)

	topic := "test_topic"

	partitionList, err := u.consumer.Partitions(topic)
	if err != nil {
		log.Fatal(err)
	}

	for _, partition := range partitionList {
		go consumeMessage(u.consumer, topic, partition, chanMessage)
	}

	for {
		msg , ok := <- chanMessage
		if !ok {
			break
		}
		log.Println(string(msg.Value))
		check := &notif.NotifRequest{}
		err = proto.Unmarshal(msg.Value, check)
		if err != nil {
			break
		}
		log.Println(check.Body, "check")
	}
}



func consumeMessage(consume sarama.Consumer, topic string, partition int32, c chan *sarama.ConsumerMessage) {
	msg, err := consume.ConsumePartition(topic, partition, sarama.OffsetNewest)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := consume.Close(); err != nil {
			log.Fatal(err)
			return
		}
	}()

	for {
		msg := <- msg.Messages()
		c <- msg
	}	

}