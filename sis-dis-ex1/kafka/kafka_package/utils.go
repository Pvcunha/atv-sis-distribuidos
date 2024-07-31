package kafka_sample

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func ReadConfig() kafka.ConfigMap {

	conf := &kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1:9092",
		"group.id":          "test-kafka",
		"auto.offset.reset": "latest",
	}

	return *conf
}

func ExitApplication() {
	fmt.Println("Gracefully shutting down application")
	time.Sleep(time.Second * 5)
	os.Exit(1)
}
