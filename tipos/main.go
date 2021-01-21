package main

import "fmt"

type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

// declaración de alias
type myAlias = course

// definiciones de tipo, hereda las propiedades de la estructura base, pero no sus metodos
type newCourse course

// nuevo tipo predeclarado
type newBool bool

func (n newBool) String() string {
	if n {
		return "TRUE"
	} else {
		return "FALSE"
	}
}

func main() {
	c := course{name: "Go"}
	c.Print()
	fmt.Printf("El tipo es: %T\n", c)

	fmt.Println("=======Declaración de alias=======")
	d := myAlias{name: "Go"}
	d.Print()
	fmt.Printf("El tipo es: %T\n", d)

	fmt.Println("=======Definición de tipo=======")
	e := newCourse{name: "Go"}
	fmt.Printf("El tipo es: %T\n", e)

	fmt.Println("=======Nuevo tipo predeclarado=======")
	var b newBool = false
	fmt.Println(b.String())

}
