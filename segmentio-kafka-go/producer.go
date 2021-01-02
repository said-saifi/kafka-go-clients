package main

import (
	"context"
	kafka "github.com/segmentio/kafka-go"
)

// topic and broker addresses
const (
	topic   = "my-kafka-topic"
	broker1 = "localhost:9092"
	broker2 = "localhost:9093"
	broker3 = "localhost:9094"
)

// returns messages to be sent to kafka
func getMessages() []string {
	return []string{"First test message!", "Second test message!", "Third test message!", "Fourth test message!"}
}

// produce messages to kafka
func produce(ctx context.Context) error{

	// creating a new writer
	// the topic can be defined on this level as well
	k := &kafka.Writer{
		Addr: kafka.TCP(broker1, broker2, broker3),
		Balancer: &kafka.LeastBytes{},  // Other types implementing Balancer interface could be used
	}

	var kms []kafka.Message
	for _, m := range getMessages(){
		km := kafka.Message{
			Topic: topic,   // This can be defined on writer level instead
			Value: []byte(m),
		}
		kms = append(kms, km)
	}

	return k.WriteMessages(ctx, kms...)

}