package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	topic = "contact-create-email"
	group = "contact-create-email"
)

func main() {

	flag.StringVar(&group, "consumer-group", "contact-create-email", "--consumer-group=contact-create-email")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     topic,
		Partition: 0,
		GroupID:   group,
		//MaxBytes:  batchSize,
	})
	for {
		time.Sleep(time.Millisecond * 100)
		message, err := reader.ReadMessage(context.TODO())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("sending email: -->>>>>>> " + string(message.Value))
	}

}
