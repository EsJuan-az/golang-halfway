package main

import "fmt"

type PrintInfo interface{
	getInfo() string
}
func printInfo(pi PrintInfo){
	fmt.Println(pi.getInfo())
}
type TemporaryEmployee struct{
	BaseEmployee
	endDate string
}
func (te *TemporaryEmployee) getInfo() string{
	return fmt.Sprintf("%d) %s (%d años de edad): Trabaja hasta %s",te.id, te.name, te.age, te.endDate)
}
func NewTemporary(id int, name string, age int, endDate string)*TemporaryEmployee{
	te := &TemporaryEmployee{}
	te.id = id
	te.name = name
	te.age = age
	te.endDate = endDate
	return te

}
type FullTimeEmployee struct{
	BaseEmployee
	taxRate int
}
func (fte *FullTimeEmployee) getInfo() string{
	return fmt.Sprintf("%d) %s (%d años de edad): Índice de impuestos %d",fte.id, fte.name, fte.age, fte.taxRate)
}
func NewFullTime(id int, name string, age, taxRate int)*FullTimeEmployee{
	fte := &FullTimeEmployee{}
	fte.id = id
	fte.name = name
	fte.age = age
	fte.taxRate = taxRate
	return fte
}
func Interface(){
	t := NewTemporary(1, "Pepe", 15, "1 de Mayo")
	ft := NewFullTime(2, "Pedro", 51, 10)
	printInfo(t)
	printInfo(ft)
}