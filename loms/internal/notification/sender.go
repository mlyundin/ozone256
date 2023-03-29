package notification

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"

	"route256/loms/pkg/model"
)

type notificationSender struct {
	producer      sarama.SyncProducer
	asyncProducer sarama.AsyncProducer
	topic         string
}

type Handler func(id string)

func NewOrderSender(producer1 sarama.SyncProducer, producer2 sarama.AsyncProducer, topic string, onSuccess, onFailed Handler) *notificationSender {
	s := &notificationSender{
		producer:      producer1,
		asyncProducer: producer2,
		topic:         topic,
	}

	// config.Producer.Return.Errors = true
	go func() {
		for e := range producer2.Errors() {
			bytes, _ := e.Msg.Key.Encode()

			onFailed(string(bytes))
			log.Println(e.Msg.Key, e.Error())
		}
	}()

	// config.Producer.Return.Successes = true
	go func() {
		for m := range producer2.Successes() {
			bytes, _ := m.Key.Encode()

			onSuccess(string(bytes))
			log.Printf("order id: %s, partition: %d, offset: %d\n", string(bytes), m.Partition, m.Offset)
		}
	}()

	return s
}

func (s *notificationSender) SendOrderStatusUpdate(orderID int64, newStatus, oldStatus model.OrderStatus) error {
	msg := &sarama.ProducerMessage{
		Topic:     s.topic,
		Partition: -1,
		Value:     sarama.StringEncoder(fmt.Sprintf(`{"order_id": %d, "new_status": %d, "old_status": %d}`, orderID, newStatus, oldStatus)),
		Key:       sarama.StringEncoder(fmt.Sprint(orderID)),
		Timestamp: time.Now(),
	}

	partition, offset, err := s.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("order id: %d, partition: %d, offset: %d", orderID, partition, offset)
	return nil
}
