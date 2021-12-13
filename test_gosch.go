package main

import (
	"fmt"
	"runtime"
)

func sing() {
	for i := 0; i < 10; i++ {
		fmt.Println("singing...")
		runtime.Gosched()
	}
}

func dance() {
	for i := 0; i < 10; i++ {
		fmt.Println("danceing...")
		runtime.Gosched()
	}
}

func main(){

	go sing()
	go dance()
	for{
		;
	}
}
