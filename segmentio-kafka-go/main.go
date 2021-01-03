package main

import (
	"context"
	"log"
	"time"
)

func main() {
	go consume()
	if err := produce(context.Background()); err != nil {
		log.Fatalf("Error producing messages to kafka: %v", err)
	}
	time.Sleep(3*time.Second)


}


