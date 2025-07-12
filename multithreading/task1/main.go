package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello world")
	}()
	wg.Wait()

}

// Вариант 1
//	func main() {
//		done := make(chan bool)
//		go func(done chan bool) {
//			fmt.Println("Hello world")
//			close(done)
//		}(done)
//		<-done
//	}
// Вариант 2
// func main() {
// 	mutex := &sync.Mutex{}
// 	go func() {
// 		mutex.Lock()
// 		defer mutex.Unlock()
// 		fmt.Println("Hello world")
// 	}()
// 	time.Sleep(100 * time.Millisecond)
// }
