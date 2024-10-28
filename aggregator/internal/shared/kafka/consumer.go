package kafka

import (
	"log"

	"context"

	"github.com/IBM/sarama"
)

type HandlerKafka interface {
	Start(ctx context.Context, brokers []string, group string, topics []string) error
	SetConsumer(consumer *ConsumerKafka)
	HandleBookingCreated(message *sarama.ConsumerMessage)
	HandleBookingBegin(message *sarama.ConsumerMessage)
	HandleBookingUpdated(message *sarama.ConsumerMessage)
	HandleBookingFinished(message *sarama.ConsumerMessage)
	HandleBookingCancelled(message *sarama.ConsumerMessage)
	HandleApartmentCreated(message *sarama.ConsumerMessage)
	HandleApartmentRemoved(message *sarama.ConsumerMessage)
	HandleApartmentUpdated(message *sarama.ConsumerMessage)
	HandleCustomerCreated(message *sarama.ConsumerMessage)
	HandleCustomerRemoved(message *sarama.ConsumerMessage)
	HandleCustomerUpdated(message *sarama.ConsumerMessage)
}

type ConsumerKafka struct {
	Ready   chan bool
	handler HandlerKafka
}

func NewConsumer(handler HandlerKafka) *ConsumerKafka {
	return &ConsumerKafka{
		Ready:   make(chan bool),
		handler: handler,
	}
}

func (consumer *ConsumerKafka) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.Ready)
	return nil
}

func (consumer *ConsumerKafka) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *ConsumerKafka) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		switch message.Topic {
		case "booking_created":
			consumer.handler.HandleBookingCreated(message)
		case "booking_begin":
			consumer.handler.HandleBookingBegin(message)
		case "booking_updated":
			consumer.handler.HandleBookingUpdated(message)
		case "booking_finished":
			consumer.handler.HandleBookingFinished(message)
		case "booking_cancelled":
			consumer.handler.HandleBookingCancelled(message)
		case "apartment_created":
			consumer.handler.HandleApartmentCreated(message)
		case "apartment_removed":
			consumer.handler.HandleApartmentRemoved(message)
		case "apartment_updated":
			consumer.handler.HandleApartmentUpdated(message)
		case "customer_created":
			consumer.handler.HandleCustomerCreated(message)
		case "customer_removed":
			consumer.handler.HandleCustomerRemoved(message)
		case "customer_updated":
			consumer.handler.HandleCustomerUpdated(message)
		default:
			log.Printf("Получено сообщение из неизвестного топика %s: %s", message.Topic, string(message.Value))
		}

		// Отмечаем сообщение как обработанное
		session.MarkMessage(message, "")
	}
	return nil
}
