package kafka

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/IBM/sarama"
)

type Consumer struct {
	ready chan bool
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		session.MarkMessage(message, "")
	}
	return nil
}

func ff() {
	brokers := "localhost:9092"
	group := "example_group"
	topics := "example_topic"

	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0

	consumer := Consumer{
		ready: make(chan bool),
	}

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, config)
	if err != nil {
		log.Fatalf("Error creating consumer group client: %v", err)
	}
	defer func(client sarama.ConsumerGroup) {
		err := client.Close()
		if err != nil {
			log.Fatalf("Error closing client: %v", err)
		}
	}(client)

	go func() {
		for {
			if err := client.Consume(ctx, strings.Split(topics, ","), &consumer); err != nil {
				log.Fatalf("Error from consumer: %v", err)
			}
			// Check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready
	log.Println("Sarama consumer up and running!...")

	<-ctx.Done()
	log.Println("Terminating: context cancelled")
}
