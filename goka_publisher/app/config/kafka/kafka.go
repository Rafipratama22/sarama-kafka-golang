package kafka

import (
	"example-goka/app/shared/variable"
	"log"

	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

type Kafka struct {
	Producer *goka.Emitter
}

var (
	topic goka.Stream = "example-stream"
)

func ConsumerConfig() *Kafka {
	emitter, err := goka.NewEmitter(variable.Kafka, topic, new(codec.Bytes))
	if err != nil {
		log.Fatalf("error creating emitter: %v", err)
	}

	return &Kafka{
		Producer: emitter,
	}
}
