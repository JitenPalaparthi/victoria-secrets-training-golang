// It is used to manage the lifecycle of concurrent operations that may depend on each other or have a time limit

//Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	//ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Millisecond * 1)
		cancel()
	}()
	ch := Gen(ctx)
	for v := range ch {
		fmt.Println(v)
	}

}

func Gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		i := 1
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- i * i
				i++
			}
		}
	}()
	return ch
}
