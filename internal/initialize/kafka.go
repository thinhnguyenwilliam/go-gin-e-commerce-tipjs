package initialize

import (
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/thinhcompany/ecommerce-ver-2/global"
)

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9192"),
		Topic:    "otp-auth-topic",
		Balancer: &kafka.LeastBytes{},
	}

	log.Println("✅ Kafka producer initialized on localhost:9192, topic: otp-auth-topic")
}

func CloseKafka() {
	if global.KafkaProducer != nil {
		if err := global.KafkaProducer.Close(); err != nil {
			log.Fatalf("❌ Failed to close Kafka producer: %v", err)
		} else {
			log.Println("✅ Kafka producer closed.")
		}
	}
}
