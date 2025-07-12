package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	input := make(chan int)
	defer cancel()
	go StartBatchProcessor(ctx, input)
	go func() {
		for i := 0; i < 20; i++ {
			input <- i
			time.Sleep(20 * time.Millisecond)
		}
	}()
	<-ctx.Done()
	fmt.Println("Main: processing stopped")
}

func StartBatchProcessor(ctx context.Context, input <-chan int) {
	maxBatchers := 5
	timeout := 2 * time.Millisecond
	timer := time.NewTimer(timeout)
	batch := make([]int, 0, maxBatchers)
	for {
		select {
		case x := <-input:
			batch = append(batch, x)
			if len(batch) == maxBatchers {
				batch = batch[:0]
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(timeout)
			}
		case <-ctx.Done():
			return
		case <-timer.C:
			if len(batch) > 0 {
				batch = batch[:0]
			}
			timer.Reset(timeout)

		}
	}

}
