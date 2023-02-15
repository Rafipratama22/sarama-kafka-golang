package container

import (
	"example-sarama/app/config/kafka"
	"example-sarama/app/usecase"
)

type Container struct {
	Event usecase.Usecase
}

func SetupContainer() Container {
	configKafka := kafka.ConsumerConfig()

	consumer := configKafka.Consumer

	event := usecase.NewUsecase(consumer)

	return Container{
		Event: event,
	}

}
