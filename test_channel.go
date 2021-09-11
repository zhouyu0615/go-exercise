package main

import (
	"fmt"
	"time"
)

func goroutine1(ch chan int) {

	for i := 1; i < 100; i++ {
		ch <- i
		if i%2 == 1 {
			fmt.Println("go1:", i)
		}g
	}
}

func goroutine2(ch chan int) {

	for i := 1; i < 100; i++ {
		<-ch
		if i%2 == 0 {
			fmt.Println("go2:", i)
		}
	}
}

func main() {
	ch := make(chan int)

	go goroutine1(ch)
	go goroutine2(ch)

	time.Sleep(time.Second * 2)
}
