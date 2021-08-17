package main

import (
	"fmt"
)

func add() func (int) int {
	n := 10
	m := 100
	fmt.Println("add().n of start = " , n)
	fmt.Println("add().m of start = " , m)
	return func (x int) int{
		n = n + x
		fmt.Println("add().x = " , x)
		fmt.Println("add().n of end = " , n)
		fmt.Println("add().m of end = " , m)
		return n
	}
}

func add2(base int) func (int) int {
	n := 10
	m := func (x int) int{
		n = n + x + base
		return n
	}
	//fmt.Println("n = " , n)
	return m
}

func main(){
	f1 := add()
	fmt.Println(f1(1))
	fmt.Println(f1(1))

	f2 := add2(10)
	fmt.Println(f2(1))
	fmt.Println(f1(1))

	
}