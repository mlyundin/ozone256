package main

import (
	"context"
	"log"
	"route256/libs/config"
	"route256/libs/kafka"
	receiver "route256/notifications/internal/kafka"
	"time"
)

func main() {

	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	// TODO wait for kafka brokers up
	time.Sleep(15 * time.Second)
	brokers := make([]string, 0, len(config.ConfigData.Kafka.Brokers))
	for _, broker := range config.ConfigData.Kafka.Brokers {
		brokers = append(brokers, broker.Url())
	}

	consumer, err := kafka.NewConsumer(brokers)
	if err != nil {
		log.Fatalln(err)
	}

	topic := config.ConfigData.Kafka.OrderStatusTopic
	handlers := map[string]receiver.HandleFunc{
		topic: func(id string, value string) {
			log.Println("Got order status update for id; ", id, " value: ", value)
		},
	}
	r := receiver.NewReciver(consumer, handlers)
	err = r.Subscribe(topic)
	if err != nil {
		log.Fatalln(err)
	}

	<-context.TODO().Done()
}
