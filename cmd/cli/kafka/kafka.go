package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

const (
	kafkaURL   = "localhost:9192" // update if your broker uses another port
	kafkaTopic = "test-topic"
)

var kafkaProducer *kafka.Writer

// Kafka writer (producer)
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// Kafka reader (consumer)
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	fmt.Println("Creating reader with brokers:", kafkaURL) // 🟢 Add this
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3,
		MaxBytes:       10e6,
		CommitInterval: time.Second,

		StartOffset: kafka.LastOffset, // consume only new messages: kafka.LastOffset
		//kafka.FirstOffset: consumers will consume everything, even older messages.
		//this will help you see if consumers are working correctly, regardless of when the message was published.
	})
}

// Struct to hold stock message
type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func newStock(msg, typeMsg string) *StockInfo {
	return &StockInfo{Message: msg, Type: typeMsg}
}

// Gin handler to publish message to Kafka
func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))

	body := map[string]any{
		"action": "action",
		"info":   s,
	}
	jsonBody, _ := json.Marshal(body)

	msg := kafka.Message{
		Key:   []byte("action"),
		Value: jsonBody,
	}

	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		fmt.Printf("❌ Failed to produce message: %v\n", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("✅ Message sent to Kafka")

	fmt.Printf("Produced message: %s\n", string(jsonBody))

	c.JSON(200, gin.H{
		"msg": "Action successfully published to Kafka",
	})
}

// Kafka consumer simulating an ATC stock buyer
func RegisterConsumerATC(id int) {
	groupID := fmt.Sprintf("consumer-group-%d", id)
	fmt.Printf("Consumer %d using broker: %s\n", id, kafkaURL) // 🟢 Add this
	reader := getKafkaReader(kafkaURL, kafkaTopic, groupID)
	defer reader.Close()

	fmt.Printf("Consumer %d listening for ATC stock messages...\n", id)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer(%d) error: %v\n", id, err)
			continue
		}

		fmt.Printf("[Consumer %d] Topic: %s | Offset: %d | %s = %s\n",
			id, m.Topic, m.Offset, string(m.Key), string(m.Value))
	}
}

//curl -X POST "http://localhost:8999/action/stock?msg=BUY-VN30&type=ATC"

func main() {
	fmt.Println("Kafka URL:", kafkaURL) // Add this!
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()

	r := gin.Default()
	r.POST("/action/stock", actionStock)

	// Start multiple consumers
	for i := 1; i <= 4; i++ {
		go RegisterConsumerATC(i)
	}

	r.Run(":8999")
}
