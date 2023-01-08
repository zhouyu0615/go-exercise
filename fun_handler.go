package main

import "fmt"

type testHandler func(a, b interface{}) interface{}

func dd(i testHandler) interface{} {
	fmt.Printf("i type: %T\n", i)
	return i(1, 2)
}

func main() {
	ee := func(x, y interface{}) interface{} {
		t, ok := x.(int)
		if ok {
			k, ok := y.(int)
			if ok {
				return t + k
			}
		}
		return nil
	}

	fmt.Printf("ee type: %T\n", ee)
	fmt.Println(dd(ee))
}
