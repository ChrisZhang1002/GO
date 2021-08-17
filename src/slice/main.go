package main

import (
	"fmt"
)

func main()  {
	var slice []int = make([]int, 5,10)
	fmt.Println(slice)
	fmt.Println(cap(slice))
}