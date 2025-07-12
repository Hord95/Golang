package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	res := 0
	go sum(ch)

	for x := range ch {
		res += x
	}
	fmt.Println(res)
}

// func main() {
// 	ch := make(chan int)
// 	go sum(ch)
// 	res := 0
// 	for {
// 		select {
// 		case num, ok := <-ch:

//				if !ok {
//					fmt.Println(res)
//					return
//				}
//				res += num
//			case <-time.After(time.Second):
//				return
//			}
//		}
//	}
func sum(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i

	}
	close(ch)
}
