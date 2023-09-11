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
outer:
	for {
		select {
		case <-timeOut:
			cancel()
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break outer
		default:
			time.Sleep(time.Millisecond * 100)
			message := "Hello World!--->" + fmt.Sprint(counter)
			err := writer.WriteMessages(ctx,
				kafka.Message{
					//Key:   []byte("Key-A"),
					Value: []byte(message),
				},
			)
			if err != nil {
				cancel()
			}
			fmt.Println(message)
			counter++
		}
	}

}
