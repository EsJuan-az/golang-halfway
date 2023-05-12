package main

import "fmt"
type Person struct{
	DNI string
	name string
	age int
}
type BaseEmployee struct{
	id int
	Person
}
func Inheritance(){
	be := BaseEmployee{}
	fmt.Printf("%+v\n", be)
}