package delivery

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/koriebruh/simply_microservice/cfg"
	"github.com/koriebruh/simply_microservice/entity"
	"gorm.io/gorm"
	"log"
	"time"
)

// CONSUMER ADA 2 CONSUME, dari payment dan shipping

func PaymentKafkaConsumer(config *cfg.Config, topicListen string, db *gorm.DB) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s",
			config.Kafka.Server, config.Kafka.Port),
		"group.id":          "group-payment",
		"auto.offset.reset": "earliest", // Mulai dari awal jika tidak ada offset
	}

	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		log.Fatal("got wrong in configuration kafka")
	}

	if err = consumer.Subscribe(topicListen, nil); err != nil {
		log.Fatal(fmt.Sprint("NotFound Topic ", topicListen))
	}

	for {
		message, err := consumer.ReadMessage(1 * time.Second)
		if err != nil {
			log.Fatal("failed to consume message")
		}

		go func(msg *kafka.Message) {
			var data PaymentConsumer
			if err2 := json.Unmarshal(msg.Value, &data); err2 != nil {
				log.Fatal("failed unmarshal")
			}

			if err2 := db.Model(&entity.Order{}).Where("id = ?", data.IdOrder).Updates(
				map[string]interface{}{
					"payment_method": data.PaymentMethod,
					"payment_status": data.PaymentStatus,
				}).Error; err2 != nil {
				log.Printf("Error updating order status from Kafka: %v", err)
			}
		}(message)

	}
}

type PaymentConsumer struct {
	IdOrder       uint                 `json:"id_order"`
	PaymentStatus entity.PaymentStatus `json:"payment_status"`
	PaymentMethod entity.PaymentMethod `json:"payment_method"`
}

func ShippingKafkaConsumer(config *cfg.Config, topicListen string, db *gorm.DB) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s",
			config.Kafka.Server, config.Kafka.Port),
		"group.id":          "group-consumer",
		"auto.offset.reset": "earliest", // Mulai dari awal jika tidak ada offset
	}

	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		log.Fatal("got wrong in configuration kafka")
	}

	if err = consumer.Subscribe(topicListen, nil); err != nil {
		log.Fatal(fmt.Sprint("NotFound Topic ", topicListen))
	}

	for {
		message, err := consumer.ReadMessage(1 * time.Second)
		if err != nil {
			log.Fatal("failed to consume message")
		}

		go func(msg *kafka.Message) {
			var data ShippingConsumer
			if err2 := json.Unmarshal(msg.Value, &data); err2 != nil {
				log.Fatal("failed unmarshal")
			}

			if err2 := db.Model(&entity.Order{}).Where("id = ?", data.IdOrder).Updates(
				map[string]interface{}{
					"shipping_status": data.ShippingStatus,
				}).Error; err2 != nil {
				log.Printf("Error updating order status from Kafka: %v", err)
			}
		}(message)

	}
}

type ShippingConsumer struct {
	IdOrder        uint                  `json:"id_order"`
	ShippingStatus entity.ShippingStatus `json:"shipping_status"`
}
