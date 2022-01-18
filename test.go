package main

import "fmt"
import "sort"

# for test

func main() {
   fmt.Println("Hello, World!")
   
   fmt.Println(1<<1, 1<<0)

   var abc="abc"
   fmt.Println(int(abc[0]))

   fmt.Println(int(abc[1]))
   a := fmt.Sprintf("char: %c_name", rune(97))
   fmt.Println(a)

continentArray := []string {"North America", "Asia", "South America", "Europe", "Africa", "Australia"}

   fmt.Println(continentArray)

   sort.Strings(continentArray)
   fmt.Println(continentArray)

   continentArray = continentArray[:0]   
   fmt.Println(continentArray)
   
   fmt.Println(len(continentArray))

   keys :=[]int{1,2,3}
   fmt.Println("keys:%v", keys)
   for _, key := range keys {
   	fmt.Println("key:",key)
   }






   
}
