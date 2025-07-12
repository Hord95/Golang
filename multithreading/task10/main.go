package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	n := 3
	channels := Split(ch, n)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	wg.Add(n)
	for i, c := range channels {
		go func(i int, c <-chan int) {
			defer wg.Done()
			for x := range c {
				fmt.Printf("Worker %d received %d\n", i, x)
			}
		}(i, c)
	}

	wg.Wait()
	time.Sleep(10 * time.Millisecond) // Даём время для вывода
}
func Split(ch chan int, n int) []chan int {

	cs := make([]chan int, 0, n)
	for i := 0; i < n; i++ {
		cs = append(cs, make(chan int))
	}
	toChannels := func(ch chan int, cs []chan int) {
		defer func(cs []chan int) {
			for _, c := range cs {
				close(c)
			}
		}(cs)
		for {
			for _, c := range cs {
				select {
				case val, ok := <-ch:
					if !ok {
						return
					}
					c <- val
				}
			}
		}
	}

	go toChannels(ch, cs)
	return cs

}
