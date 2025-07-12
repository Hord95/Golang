package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем 3 канала
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Запускаем горутины для заполнения каналов
	go func() {
		defer close(ch1)
		for i := 1; i <= 3; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 4; i <= 6; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer close(ch3)
		for i := 7; i <= 9; i++ {
			ch3 <- i
		}
	}()

	// Объединяем каналы
	merged := Merge(ch1, ch2, ch3)

	// Читаем из объединенного канала
	for num := range merged {
		fmt.Println(num)
	}

}
func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	send := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}

	}
	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
