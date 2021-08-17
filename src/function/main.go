package main

import "fmt"

func main(){

	res1 := func (n1 int, n2 int) int{
		return n1 + n2
	}(10,20)

	fmt.Printf("res1 = %v\n", res1)

	a := func (n1 int, n2 int) int {
		return n1 - n2
	}

	res2 := a(50 ,20)
	fmt.Printf("res2 = %d\n", res2)

	res3 := a(100,50)
	fmt.Printf("res3 = %v \n", res3)

}


func init(){
	fmt.Println("This is start from init() function....")
}