package main

import "fmt"
func Generator(c chan int){
	for i := 1; i <= 10; i++{
		c <- i
	}
	close(c)
}
func Double(c chan int, d chan int){
	for value := range c{
		d <- 2 * value
	}
	close(d)
}
func Print(c chan int){
	for value := range c{
		fmt.Println(value)
	}
}
func PipeLines(){
	generator := make(chan int)
	doubles	:= make(chan int)
	
	go Generator(generator)
	go Double(generator, doubles)
	Print(doubles)
}