package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 3 {
		i++
		fmt.Println(i)
	}

	for j := 3; j < 5; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Print("loop")
		break
	}

	for k := 1; k < 10; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Println(k)
	}

	nums := []int{1, 2, 4, 5, 6, 7}
	for k, v := range nums {
		fmt.Printf("%d -> %d\n", k, v)
	}

}
