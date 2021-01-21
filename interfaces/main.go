package main

import "fmt"

//Greeter interface
type Greeter interface {
	Greet()
}

//Byer interface
type Byer interface {
	Bye()
}

//Person struct
type Person struct {
	Name string
}

//Greet method of People
func (p Person) Greet() {
	fmt.Printf("Hola soy %s\r\n", p.Name)
}

//Bye method of People
func (p Person) Bye() {
	fmt.Printf("Chao soy %s\r\n", p.Name)
}

func (p Person) String() string {
	return p.Name
}

//Text type
type Text string

//Greet method of Text type
func (t Text) Greet() {
	fmt.Printf("Hola soy %s\r\n", t)
}

//Bye method of Text type
func (t Text) Bye() {
	fmt.Printf("Chao soy %s\r\n", t)
}

func GreetAll(gs ...Greeter) {
	for _, g := range gs {
		g.Greet()
		fmt.Printf("\t Valor: %v, Tipo: %T\r\n", g, g)
	}
}

func ByeAll(gs ...Byer) {
	for _, g := range gs {
		g.Bye()
		fmt.Printf("\t Valor: %v, Tipo: %T\r\n", g, g)
	}
}

//Tipos embebidos (composicion de interfaces)
type GreeterByer interface {
	Greeter
	Byer
}

func All(gs ...GreeterByer) {
	for _, gb := range gs {
		gb.Greet()
		gb.Bye()
	}
}

func main() {
	g := Person{Name: "César"}

	// Imprime el nombre desde la interfaz definida
	fmt.Println(g)

	// g.Greet()
	// t := Text("Daisy")
	// t.Greet()

	// GreetAll(g, t)

	// GreetAll(g, t)
	// ByeAll(g, t)
}
