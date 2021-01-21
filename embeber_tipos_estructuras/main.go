package main

import "fmt"

//Person struct
type Person struct {
	Name string
	Age  uint
}

//NewPerson construct
func NewPerson(name string, age uint) Person {
	return Person{Name: name, Age: age}
}

//Greet method of Person structure
func (p Person) Greet() {
	fmt.Println("Hello from person")
}

//Human struct
type Human struct {
	Age      uint
	Children uint
}

//NewHuman construct
func NewHuman(age, children uint) Human {
	return Human{Age: age, Children: children}
}

//Employee struct
type Employee struct {
	Person // hereda los campos y metodos de Person para Employee, pero cuando se llamen, el receptor será Person
	Human
	Salary float64
}

//NewEmployee constructor
func NewEmployee(name string, age uint, children uint, salary float64) Employee {
	return Employee{NewPerson(name, age), NewHuman(age, children), salary}
}

//Payroll method of Employee
func (e Employee) Payroll() {
	fmt.Println(e.Salary * 0.90)
}

//Greet method of Person structure // ocultacion de metodos, no sobreescritura de metodos
func (e Employee) Greet() {
	fmt.Println("Hello from employee")
}

func main() {
	e := NewEmployee("César", 30, 2, 9000000)
	fmt.Println(e.Name, e.Human.Age)
	e.Greet()
	e.Person.Greet()
	e.Payroll()
}
