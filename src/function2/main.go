package main

import "fmt"

func main(){

	defer fmt.Println("it should be 2nd ")
	fmt.Println("it should be 1st")
	sub1()
}

func sub1(){
	defer fmt.Println("in sub1 . 2nd ")
	fmt.Println("in sub1 . 1st")
}