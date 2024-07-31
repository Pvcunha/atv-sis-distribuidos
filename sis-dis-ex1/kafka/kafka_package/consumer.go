package kafka_sample

import (
	"errors"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Consume(topic string, run *bool) (result chan string, err error) {

	channel := make(chan string)

	conf := ReadConfig()
	c, err := kafka.NewConsumer(&conf)

	if err != nil {
		close(channel)
		return channel, errors.New(fmt.Sprintf("[CONSUMER] Failed to create consumer: %s", err))
	}

	if err = c.SubscribeTopics([]string{topic}, nil); err != nil {
		close(channel)
		return channel, errors.New(fmt.Sprintf("[CONSUMER] Failed to subscribe topics: %s", err))
	}

	go func() {
		for *run {
			if ev, err := c.ReadMessage(200 * time.Millisecond); err != nil {
				continue
			} else {
				channel <- fmt.Sprintf("[CONSUMER] Consumed event from topic %s: value = %s",
					*ev.TopicPartition.Topic, string(ev.Value))
			}
		}

		if err = c.Close(); err != nil {
			channel <- fmt.Sprintf("[CONSUMER] Failed to close consumer: %s", err)
		}

		channel <- fmt.Sprintf("[CONSUMER] Stopping consumer")
	}()

	return channel, nil
}
