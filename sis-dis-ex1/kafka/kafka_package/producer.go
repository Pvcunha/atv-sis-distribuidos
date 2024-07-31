package kafka_sample

import (
	"errors"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Produce(topic string, run *bool) (result <-chan string, err error) {

	channel := make(chan string)

	conf := ReadConfig()
	p, err := kafka.NewProducer(&conf)

	if err != nil {
		return channel, errors.New(fmt.Sprintf("[PRODUCER] Failed to create producer: %s", err))
	}

	// Go-routine to handle message delivery reports and possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					channel <- fmt.Sprintf("[PRODUCER] Failed to deliver message: %v", ev.TopicPartition)
				} else {
					channel <- fmt.Sprintf("[PRODUCER] Produced event to topic %s: value = %s", *ev.TopicPartition.Topic, string(ev.Value))
				}
			}
		}
	}()

	go func() {
		i := 0
		for *run {
			message := &kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Key:            nil,
				Value:          []byte(fmt.Sprintf("Message nº %d", i)),
			}
			if err = p.Produce(message, nil); err != nil {
				channel <- fmt.Sprintf("[PRODUCER] Error producing message")
			}
			i++
			time.Sleep(time.Second * 2)
		}
		channel <- fmt.Sprintf("[PRODUCER] Stopping sending messages")
		p.Close()
	}()

	return channel, nil
}
