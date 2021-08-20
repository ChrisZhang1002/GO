package main

import (
	"fmt"
)

type people struct{
	Name string
	Score int
}

type Student struct{
	people
	class string
}


func main(){

	var a Student

	a.Name = "abc"
	a.Score = 123
	a.class = "class 3"

	fmt.Println(a)
	
}