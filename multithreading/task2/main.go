package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go num(&wg, &mu)

	}
	wg.Wait()

}

func num(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for i := 1; i < 6; i++ {
		mu.Lock()
		fmt.Println(i)
		mu.Unlock()
	}

}

//Вариант 2
// func main() {
// 	ch := make(chan bool)
// 	for i := 1; i <= 5; i++ {
// 		go func() {
// 			for j := 1; j <= 5; j++ {
// 				<-ch
// 				fmt.Println(j)
// 			}
// 		}()

// 		for i := 1; i <= 5; i++ {
// 			ch <- true
// 		}
// 	}
// }
//Вариант 3
// func main() {
// 	for i := 0; i < 5; i++ {
// 		go func() {
// 			for j := 1; j <= 5; j++ {
// 				fmt.Println(j)
// 			}
// 		}()
// 		time.Sleep(10 * time.Millisecond)

// 	}

// }
