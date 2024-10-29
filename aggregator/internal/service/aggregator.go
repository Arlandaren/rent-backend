package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"service/internal/shared/kafka"
)

type Aggregator struct {
	service  *Service
	consumer *kafka.ConsumerKafka
}

func NewAggregator(svc *Service, consumer *kafka.ConsumerKafka) *Aggregator {
	return &Aggregator{
		service:  svc,
		consumer: consumer,
	}
}

func (aggregator *Aggregator) SetConsumer(consumer *kafka.ConsumerKafka) {
	aggregator.consumer = consumer
}

func (aggregator *Aggregator) Start(ctx context.Context, brokers []string, group string, topics []string) error {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{
		sarama.NewBalanceStrategyRange(),
	}

	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	client, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		return fmt.Errorf("error creating consumer group: %v", err)
	}

	// Создаем канал для ошибок
	consumerErrors := make(chan error, 1)

	// Запускаем в отдельной горутине цикл потребления
	go func() {
		defer close(consumerErrors)
		for {
			// Проверяем контекст на отмену
			if ctx.Err() != nil {
				return
			}
			// Начинаем потребление
			err := client.Consume(ctx, topics, aggregator.consumer)
			if err != nil {
				consumerErrors <- err
				return
			}
			// Сбрасываем готовность после ребалансировки
			aggregator.consumer.Ready = make(chan bool)
		}
	}()

	// Ожидаем, пока потребитель будет готов
	<-aggregator.consumer.Ready
	log.Println("ConsumerKafka is ready and waiting for messages...")

	// Ожидаем завершения работы или возникновения ошибки
	select {
	case <-ctx.Done():
		log.Println("Context canceled. Exiting ConsumerKafka...")
		return client.Close()
	case err := <-consumerErrors:
		return fmt.Errorf("error in ConsumerKafka: %v", err)
	}
}

func (aggregator *Aggregator) HandleBookingCreated(message *sarama.ConsumerMessage) {
	var data BookingCreatedEvent
	err := json.Unmarshal(message.Value, &data)
	if err != nil {
		log.Printf("decode error booking_created: %v", err)
		return
	}

	err = aggregator.service.ProcessBookingCreated(&data)

	if err != nil {
		log.Printf("handle error booking_created: %v", err)
	}
}

func (aggregator *Aggregator) HandleBookingBegin(message *sarama.ConsumerMessage) {
	fmt.Printf("handling booking_begin: %s\n", string(message.Value))
	// Здесь добавь
}

func (aggregator *Aggregator) HandleBookingUpdated(message *sarama.ConsumerMessage) {
	fmt.Printf("handling booking_updated: %s\n", string(message.Value))
	// Здесь добавь
}

func (aggregator *Aggregator) HandleBookingFinished(message *sarama.ConsumerMessage) {
	fmt.Printf("handling booking_finished: %s\n", string(message.Value))
	// Здесь добавь
}

func (aggregator *Aggregator) HandleBookingCancelled(message *sarama.ConsumerMessage) {
	fmt.Printf("handling booking_cancelled: %s\n", string(message.Value))
	// Здесь добавь
}

func (aggregator *Aggregator) HandleApartmentCreated(message *sarama.ConsumerMessage) {
	fmt.Printf("handling apartment_created: %s\n", string(message.Value))
	// Здесь добавь
}

func (aggregator *Aggregator) HandleApartmentRemoved(message *sarama.ConsumerMessage) {
	fmt.Printf("handling apartment_removed: %s\n", string(message.Value))
	// Здесь добавь
}

func (aggregator *Aggregator) HandleApartmentUpdated(message *sarama.ConsumerMessage) {
	fmt.Printf("handling apartment_updated: %s\n", string(message.Value))
	// Здесь добавь
}

func (aggregator *Aggregator) HandleCustomerCreated(message *sarama.ConsumerMessage) {
	fmt.Printf("handling customer_created: %s\n", string(message.Value))
	// Здесь добавь
}

func (aggregator *Aggregator) HandleCustomerRemoved(message *sarama.ConsumerMessage) {
	fmt.Printf("handling customer_removed: %s\n", string(message.Value))
	// Здесь добавь
}

func (aggregator *Aggregator) HandleCustomerUpdated(message *sarama.ConsumerMessage) {
	fmt.Printf("handling customer_updated: %s\n", string(message.Value))
	// Здесь добавь
}
