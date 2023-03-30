package receiver

import (
	"errors"
	"log"

	"github.com/Shopify/sarama"
)

type HandleFunc func(id string, value []byte)

type Reciver struct {
	consumer sarama.Consumer
	handlers map[string]HandleFunc
}

func NewReciver(consumer sarama.Consumer, handlers map[string]HandleFunc) *Reciver {
	return &Reciver{
		consumer: consumer,
		handlers: handlers,
	}
}

func (r *Reciver) Subscribe(topic string) error {
	handler, ok := r.handlers[topic]
	if !ok {
		return errors.New("no handler for topic")
	}

	partitionList, err := r.consumer.Partitions(topic) //get all partitions on the given topic
	if err != nil {
		return err
	}

	for _, partition := range partitionList {

		initialOffset := sarama.OffsetNewest // есть риск потерять сообщения
		// initialOffset := sarama.OffsetOldest // перечитываете одни и теже сообщения
		//initialOffset := offsets[partition] // Получаем оффсет последний из внешнего storage(хранилища/БД/кеша)

		pc, err := r.consumer.ConsumePartition(topic, partition, initialOffset)
		if err != nil {
			return err
		}

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				k := string(message.Key)
				handler(k, message.Value)
				log.Printf("read: key: %s, topic: %s, partion: %d, offset: %d",
					k, topic, message.Partition, message.Offset)
			}
		}(pc)
	}

	return nil
}
