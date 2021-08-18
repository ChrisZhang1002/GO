package main

import (
	"fmt"
)

func main(){

	//var a map[string]string
	var a = make (map[string]string, 10)
	a["1"] = "first"
	a["2"] = "second"
	//fmt.Println(a)

	b := make(map[string]string)
	b["1"]="aaa"
	b["2"]="bbb"
	b["3"]="ccc"
	b["4"]="ddddd"
	b["5"]="eeeeeeee"
	b["6"]="ffffffffff"
	//fmt.Println(b)

	for k,v := range b {
		fmt.Printf("k : %s , v : %s \n", k , v)
	}



}