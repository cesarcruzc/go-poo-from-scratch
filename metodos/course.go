package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func (c *Course) ChangePrice(price float64) {
	c.Price = price
}

//PrintClasses Metodo que imprime los cursos
func (c Course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += fmt.Sprintf("%s, ", class)
	}

	fmt.Println(text[:len(text)-2])
}
