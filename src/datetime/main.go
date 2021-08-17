package main

import "fmt"

func main()  {
	a := new(int)
	fmt.Printf("a type is %T , value is %v , address is %v", a , *a , &a)
}