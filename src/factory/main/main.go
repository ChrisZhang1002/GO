package main

import (
		"fmt"
		"factory/model"
	)


func main(){

	//var stu = model.Student{
	//	Name : "abc",
	//	Score : 123,
	//}

	stu1 := model.SetStu("haha",456)
	var stu2 model.Student

	//fmt.Println("stu : ",stu)
	fmt.Println("stu1 : ", stu1)
	fmt.Println("stu2 : ", stu2.Setst("lala",123))
}