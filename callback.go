package main

import "fmt"

func Callback(){
	x := 5
	y := func(a int)int{
		return a * 2
	}
	fmt.Println(y(x))
}