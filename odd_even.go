package main

import (
	"fmt"
	"time"
)

const MAX = 10

func go1(p chan int) {
	for i := 0; i < MAX; i++ {
		p <- i
		if i%2 == 0 {
			fmt.Println("go1:", i)
		}
	}
}

func go2(p chan int) {

	for i := 0; i < MAX; i++ {
		<-p
		if i%2 == 1 {
			fmt.Println("go2", i)
		}
	}

}

func main() {

	msg := make(chan int)

	go go1(msg)
	go go2(msg)

	time.Sleep(time.Second * 1)

}
