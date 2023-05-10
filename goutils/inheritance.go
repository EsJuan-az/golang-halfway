package goutils

import "fmt"
type Person struct{
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