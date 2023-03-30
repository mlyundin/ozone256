package main

import (
	"context"
	"log"
	"route256/libs/config"
	"route256/libs/kafka"
	desc "route256/loms/pkg/loms"
	receiver "route256/notifications/internal/kafka"
	"time"

	"google.golang.org/protobuf/proto"
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
		topic: func(id string, value []byte) {
			status := &desc.OrderUpdateNotification{}
			err := proto.Unmarshal(value, status)
			if err != nil {
				log.Println("Failed to unmarshal order status notification:", err)
			} else {
				log.Println("For order: ", status.GetOrderId(), " status update: ",
					desc.OrderStatus_name[int32(status.GetOldStatus())], " -> ", desc.OrderStatus_name[int32(status.GetNewStatus())])
			}
		},
	}
	r := receiver.NewReciver(consumer, handlers)
	err = r.Subscribe(topic)
	if err != nil {
		log.Fatalln(err)
	}

	<-context.TODO().Done()
}
