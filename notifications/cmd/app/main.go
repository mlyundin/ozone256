package main

import (
	"context"
	"fmt"
	"log"
	"route256/libs/config"
	"route256/libs/kafka"
	"route256/libs/logger"
	desc "route256/loms/pkg/loms"
	receiver "route256/notifications/internal/kafka"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func main() {

	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	logger.Init(config.ConfigData.Services.Logging.Devel)

	brokers := make([]string, 0, len(config.ConfigData.Kafka.Brokers))
	for _, broker := range config.ConfigData.Kafka.Brokers {
		brokers = append(brokers, broker.Url())
	}

	consumer, err := kafka.NewConsumer(brokers)
	if err != nil {
		logger.Fatal("failed to create kafka consumer", zap.Error(err))
	}

	topic := config.ConfigData.Kafka.OrderStatusTopic
	handlers := map[string]receiver.HandleFunc{
		topic: func(id string, value []byte) {
			status := &desc.OrderUpdateNotification{}
			err := proto.Unmarshal(value, status)
			if err != nil {
				logger.Error("Failed to unmarshal order status notification:", zap.Error(err))
			} else {
				logger.Info(fmt.Sprint("For order: ", status.GetOrderId(), " status update: ",
					desc.OrderStatus_name[int32(status.GetOldStatus())], " -> ", desc.OrderStatus_name[int32(status.GetNewStatus())]))
			}
		},
	}
	r := receiver.NewReciver(consumer, handlers)
	err = r.Subscribe(topic)
	if err != nil {
		logger.Fatal("failed to subscribe", zap.Error(err))
	}

	<-context.TODO().Done()
}
