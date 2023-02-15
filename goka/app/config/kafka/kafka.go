package kafka

import (
	"log"

	"github.com/Shopify/sarama"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

var (
	brokers             = []string{"localhost:9092"}
	topic   goka.Stream = "example-stream"
	group   goka.Group  = "example-group"

	tmc *goka.TopicManagerConfig
)

type Kafka struct {
	Consumer *goka.Processor
}

func ConsumerConfig(in func(ctx goka.Context, msg interface{})) *Kafka {
	config := goka.DefaultConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	goka.ReplaceGlobalConfig(config)
	tmc := goka.NewTopicManagerConfig()
	tmc.Table.Replication = 1
	tmc.Stream.Replication = 1

	tm, err := goka.NewTopicManager(brokers, goka.DefaultConfig(), tmc)
	if err != nil {
		log.Fatalf("Error creating topic manager: %v", err)
	}
	defer tm.Close()
	err = tm.EnsureStreamExists(string(topic), 8)
	if err != nil {
		log.Printf("Error creating kafka topic %s: %v", topic, err)
	}

	g := goka.DefineGroup(group,
		goka.Input(topic, new(codec.Bytes), in),
		goka.Persist(new(codec.Int64)),
	)

	p, err := goka.NewProcessor(brokers,
		g,
		goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(tmc)),
		goka.WithConsumerGroupBuilder(goka.DefaultConsumerGroupBuilder),
	)
	if err != nil {
		log.Fatalf("error creating processor: %v", err)
	}
	log.Println("check")
	// ctx := context.Background()
	// done := make(chan struct{})
	// go func() {
	// 	defer close(done)
	// 	if err = p.Run(ctx); err != nil {
	// 		log.Printf("error running processor: %v", err)
	// 	}
	// }()

	return &Kafka{
		Consumer: p,
	}
}
