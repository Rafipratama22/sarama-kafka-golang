package container

import (
	"context"
	"example-goka/app/config/kafka"
	"example-goka/app/usecase"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Container struct {
	Event usecase.Usecase
}

func SetupContainer() Container {
	event := usecase.NewUsecase()
	eventKafka, _ := event.Notif()
	consumer := kafka.ConsumerConfig(eventKafka)
	done := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer close(done)
		err := consumer.Consumer.Run(ctx)
		if err != nil {
			log.Printf("error running processor: %v", err)
		}
	}()

	sigs := make(chan os.Signal)
	go func() {
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	}()

	select {
	case <-sigs:
	case <-done:
	}
	cancel()
	<-done

	return Container{
		Event: event,
	}

}
