package main

import (
	"fmt"
	"time"
)

func A(i time.Duration, c chan<- int, param int){
	time.Sleep(i)
	c <- param
}
func Multiplex(){
	c1 := make(chan int)
	c2 := make(chan int)
	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go A(d1, c1, 1)
	go A(d2, c2, 2)

	for i:=0; i < 2; i++{
		select{
		case msg1 := <- c1:
			fmt.Println(msg1)
		case msg2 := <- c2:
			fmt.Println(msg2)
		}
	}

}