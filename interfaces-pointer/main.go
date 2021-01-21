package main

import "fmt"

//Storager interface
type Storager interface {
	Get() string
	Set(string)
}

type Person struct {
	name []string
}

func NewPerson() *Person {

	init := []string{}
	init = append(init, "Initial")

	return &Person{
		name: init,
	}
}

func (p *Person) Get() []string {
	return p.name
}

func (p *Person) Set(name string) {
	p.name = append(p.name, name)
}

// func Exec(s Storager, name string) {
// 	s.Set(name)
// 	fmt.Println(s.Get())
// }

func main() {
	p := NewPerson()
	// fmt.Println(p.Get())
	p.Set("Camilo")
	fmt.Println(p.Get())

	// Exec(p, "Andrea")
}
