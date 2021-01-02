package main

import (
	"context"
	"log"
)

func main() {
	if err := produce(context.Background()); err != nil {
		log.Fatalf("Error producing messages to kafka: %v", err)
	}
}


