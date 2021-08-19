package main

import (
	"fmt"
)

type Person struct{
	Name string
	age int
	sex string
	slice []int
	score map[string]string
}



func main()  {

	var p1 Person 

	p1.Name = "abc"
	p1.age = 18
	p1.sex = "male"
	p1.slice = make([]int,1)
	p1.slice[0] = 100
	p1.score = make(map[string]string)
    p1.score["1"] = "abc"

	fmt.Println("p1 info : ", p1)

	p2 := Person{}

	fmt.Println(p2)


}