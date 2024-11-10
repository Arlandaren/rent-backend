package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"service/internal/shared/entities"
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
	var data entities.BookingCreatedEvent

	err := json.Unmarshal(message.Value, &data)
	if err != nil {
		log.Printf("decode error booking_created: %v", err)
		return
	}

	ctx := context.Background()

	err = aggregator.service.ProcessBookingCreated(ctx, &data)

	if err != nil {
		log.Printf("handle error booking_created: %v", err)
	}
}

func (aggregator *Aggregator) HandleBookingBegin(message *sarama.ConsumerMessage) {
	fmt.Printf("handling booking_begin: %aggregator\n", string(message.Value))

	var data entities.BookingBeganEvent

	err := json.Unmarshal(message.Value, &data)
	if err != nil {
		log.Printf("decode error booking_begin: %v", err)
		return
	}

	ctx := context.Background()

	err = aggregator.service.ProcessBookingBegan(ctx, &data)

	if err != nil {
		log.Printf("handle error booking_begin: %v", err)
	}
	log.Printf("booking_begin: %v", string(message.Value))
}

func (aggregator *Aggregator) HandleBookingUpdated(message *sarama.ConsumerMessage) {
	fmt.Printf("handling booking_updated: %aggregator\n", string(message.Value))

	var data entities.BookingUpdatedEvent

	err := json.Unmarshal(message.Value, &data)
	if err != nil {
		log.Printf("decode error booking_updated: %v", err)
		return
	}

	ctx := context.Background()

	err = aggregator.service.ProcessBookingUpdated(ctx, &data)

	if err != nil {
		log.Printf("handle error booking_updated: %v", err)
	}
	log.Printf("booking_updated: %v", string(message.Value))
}

func (aggregator *Aggregator) HandleBookingFinished(message *sarama.ConsumerMessage) {
	fmt.Printf("handling booking_finished: %aggregator\n", string(message.Value))

	var data entities.BookingFinishedEvent

	err := json.Unmarshal(message.Value, &data)
	if err != nil {
		log.Printf("decode error booking_finished: %v", err)
		return
	}

	ctx := context.Background()

	err = aggregator.service.ProcessBookingFinished(ctx, &data)

	if err != nil {
		log.Printf("handle error booking_finished: %v", err)
	}
	log.Printf("booking_finished: %v", string(message.Value))
}

func (aggregator *Aggregator) HandleBookingCancelled(message *sarama.ConsumerMessage) {
	fmt.Printf("handling booking_cancelled: %aggregator\n", string(message.Value))

	var data entities.BookingCancelledEvent

	err := json.Unmarshal(message.Value, &data)
	if err != nil {
		log.Printf("decode error booking_cancelled: %v", err)
		return
	}

	ctx := context.Background()

	err = aggregator.service.ProcessBookingCancelled(ctx, &data)

	if err != nil {
		log.Printf("handle error booking_cancelled: %v", err)
	}
	log.Printf("booking_cancelled: %v", string(message.Value))
}

func (aggregator *Aggregator) HandleApartmentCreated(message *sarama.ConsumerMessage) {
	fmt.Printf("handling apartment_created: %aggregator\n", string(message.Value))

	var data entities.ApartmentCreatedEvent

	err := json.Unmarshal(message.Value, &data)
	if err != nil {
		log.Printf("decode error apartment_created: %v", err)
		return
	}

	ctx := context.Background()

	err = aggregator.service.ProcessApartmentCreated(ctx, &data)

	if err != nil {
		fmt.Println("handle error apartment_created: ", err.Error())
		return
	}

	fmt.Println("apartment_created: ", string(message.Value))

}

func (aggregator *Aggregator) HandleApartmentRemoved(message *sarama.ConsumerMessage) {
	fmt.Printf("handling apartment_removed: %aggregator\n", string(message.Value))

	var data entities.ApartmentRemovedEvent

	err := json.Unmarshal(message.Value, &data)
	if err != nil {
		log.Printf("decode error apartment_removed: %v", err)
		return
	}

	ctx := context.Background()

	err = aggregator.service.ProcessApartmentRemoved(ctx, &data)

	if err != nil {
		fmt.Println("handle error apartment_removed: ", err.Error())
		return
	}

	fmt.Println("apartment_removed: ", string(message.Value))

}

func (aggregator *Aggregator) HandleApartmentUpdated(message *sarama.ConsumerMessage) {
	fmt.Printf("handling apartment_updated: %aggregator\n", string(message.Value))

	var data entities.ApartmentUpdatedEvent

	err := json.Unmarshal(message.Value, &data)
	if err != nil {
		log.Printf("decode error apartment_updated: %v", err)
		return
	}

	ctx := context.Background()

	err = aggregator.service.ProcessApartmentUpdated(ctx, &data)

	if err != nil {
		fmt.Println("handle error apartment_updated: ", err.Error())
		return
	}

	fmt.Println("apartment_updated: ", string(message.Value))
}

func (aggregator *Aggregator) HandleCustomerCreated(message *sarama.ConsumerMessage) {
	log.Printf("handling customer_created: %aggregator\n", string(message.Value))

	var data entities.CustomerCreatedEvent

	err := json.Unmarshal(message.Value, &data)
	if err != nil {
		log.Printf("decode error customer_created: %v", err)
		return
	}

	ctx := context.Background()

	err = aggregator.service.ProcessCustomerCreated(ctx, &data)

	if err != nil {
		log.Println("handle error customer_created: ", err.Error())
		return
	}
	log.Println("customer_created: ", string(message.Value))
}

func (aggregator *Aggregator) HandleCustomerRemoved(message *sarama.ConsumerMessage) {
	fmt.Printf("handling customer_removed: %aggregator\n", string(message.Value))

	var data entities.CustomerRemovedEvent

	err := json.Unmarshal(message.Value, &data)
	if err != nil {
		log.Printf("decode error customer_removed: %v", err)
		return
	}

	ctx := context.Background()

	err = aggregator.service.ProcessCustomerRemoved(ctx, &data)

	if err != nil {
		fmt.Println("handle error customer_removed: ", err.Error())
		return
	}

	fmt.Println("customer_removed: ", string(message.Value))
}

func (aggregator *Aggregator) HandleCustomerUpdated(message *sarama.ConsumerMessage) {
	fmt.Printf("handling customer_updated: %aggregator\n", string(message.Value))

	var data entities.CustomerUpdatedEvent

	err := json.Unmarshal(message.Value, &data)
	if err != nil {
		log.Printf("decode error customer_updated: %v", err)
		return
	}

	ctx := context.Background()

	err = aggregator.service.ProcessCustomerUpdated(ctx, &data)

	if err != nil {
		fmt.Println("handle error customer_updated: ", err.Error())
		return
	}

	fmt.Println("customer_updated: ", string(message.Value))

}
