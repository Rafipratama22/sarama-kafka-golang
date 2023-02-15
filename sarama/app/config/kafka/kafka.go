package kafka

import (
	"example-sarama/app/shared/variable"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

type Kafka struct {
	Consumer sarama.Consumer
}

func ConsumerConfig() *Kafka {

	username := variable.UsernameKafka
	password := variable.PasswordKafka

	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	
	consumers, err := sarama.NewConsumer(variable.Kafka, kafkaConfig)

	if err != nil {
		log.Fatal(err)
	}

	return &Kafka{
		Consumer: consumers,
	}
}
