package goutils

import "fmt"
type Employee struct{
	id int
	name string
}
func (e *Employee) SetId(id int){
	e.id = id
}
func (e *Employee) SetName(name string){
	e.name = name
}
func NewEmployee(id int, name string)*Employee{
	return &Employee{
		id: id,
		name: name,
	}
}
func Class(){
	e := Employee{}

	fmt.Printf("%+v\n", e)

	e.id = 12
	e.name = "Juan Carlos"
	fmt.Printf("%+v\n", e)

	e.SetId(14)
	e.SetName("Juanes teban")
	fmt.Printf("%+v\n", e)
	 
}