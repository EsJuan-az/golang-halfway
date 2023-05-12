package main

import (
	"fmt"
	"strconv"
	"time"
)
func hacerAlgo(c chan int){
	time.Sleep(5 * time.Millisecond)
	fmt.Println("Hice algo")
	c <- 1
}
func Review(){
	var x int = 8
	y := 7
	fmt.Println(x, y)
	result, err := strconv.Atoi("7890")
	if err != nil{
		fmt.Printf("%v\n", err)
	}else{
		fmt.Printf("Mi valor: %d\n", result)
	}

	edades := make(map[string]int)
	edades["Juanes"] = 17
	edades["Vane"] = 17
	edades["Sebitas"] = 17
	edades["Emita El Sabio"] = 500 //Muy sabio
	edades["Samuel"] = 16
	fmt.Printf("%#v\n", edades)
	fmt.Println()
	estudiantes := []string{"Vane", "Sebitas", "Emita El Sabio", "Samuel"}
	estudiantes = append(estudiantes, "Vin")
	for _, estudiante := range estudiantes{
		fmt.Printf("%s tiene casi %d añitos :3\n", estudiante, edades[estudiante])
	}
	c := make(chan int)
	go hacerAlgo(c)
	fmt.Println(<-c)

	g := 5832
	h := &g
	f := *h
	
	fmt.Println(g, h, f)

}