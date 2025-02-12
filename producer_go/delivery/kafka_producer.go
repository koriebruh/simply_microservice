package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/koriebruh/simply_microservice/cfg"
)

// OrderKafkaProducer message bebas si bisa struct ot apa soalnya decode to byte nya di alamnya automatic
func OrderKafkaProducer(config *cfg.Config, topicPublish string, message any) error {
	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s",
			config.Kafka.Server, config.Kafka.Port),
	}

	producer, err := kafka.NewProducer(kafkaConfig)
	if err != nil {
		_ = errors.New("got wrong in configuration kafka")
	}

	//CONVERT TO BYTE
	msgMarshal, _ := json.Marshal(message)

	sendMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topicPublish,
			Partition: kafka.PartitionAny,
		},
		Value: msgMarshal,
	}

	if err = producer.Produce(sendMessage, nil); err != nil {
		_ = errors.New("error to publish to kafka")
	}

	producer.Flush(5 * 1000)

	return nil
}
