package main

import(
	"fmt"
)

type A struct{
	num int
	aa map[string]string
}

func (a A) test()  {
	fmt.Println("a.num : ",a.num)
}

func (b A) test1(x int)  {
	b.aa = make(map[string]string)
	b.aa["1"]="first"
	b.aa["2"]="second"
	fmt.Println("x :",x)
	fmt.Println("map :",b)
	b.num = x
	fmt.Println("num should be ", b.num)
}

func main()  {
	
	var p  A
	//p.test()
	p.test1(3)
    //fmt.Println(p.num)
}