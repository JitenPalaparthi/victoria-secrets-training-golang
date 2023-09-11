package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	topic = "contact-create"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	timeOut := time.After(time.Second * 30)
	writer := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	defer writer.Close()
	counter := 1
	for {
		select {
		case <-timeOut:
			cancel()
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		default:
			time.Sleep(time.Millisecond * 100)
			err := writer.WriteMessages(ctx,
				kafka.Message{
					//Key:   []byte("Key-A"),
					Value: []byte("Hello World!--->" + fmt.Sprint(counter)),
				},
			)
			if err != nil {
				cancel()
			}
			counter++
		}
	}

}
