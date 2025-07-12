package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	res := 0
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			res++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(res)
}
