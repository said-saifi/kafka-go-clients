package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

const(
	groupID = "service-1"
)

func consume() error{
	// make a new reader that consumes from topic, partition 0
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker1, broker2, broker3},
		Topic:     topic,
		//Partition: 0,
		GroupID:  groupID,
		//GroupBalancers: []kafka.GroupBalancer{kafka.RangeGroupBalancer{}, kafka.RoundRobinGroupBalancer{}},
		MinBytes:  0,    // 0
		MaxBytes:  10e6, // 10MB

	})
	// r.SetOffset(42) // ignore first messages with offset < 42

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			return err
		}

		log.Printf("message at offset %d and paritition:%d: %s = %s\n", m.Offset, m.Partition, string(m.Key), string(m.Value))
	}

	return r.Close()

}