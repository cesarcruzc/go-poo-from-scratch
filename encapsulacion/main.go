package main

import (
	"fmt"
)

func main() {
	Go := &Course{
		Name:    "Curso de GO",
		Price:   12.34,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	Bd := Course{
		"Curso de BD",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	css := Course{
		Name:   "Curso de CSS",
		IsFree: false,
	}

	js := Course{}
	js.Name = "Curso JS"
	js.UserIDs = []uint{12, 56}

	fmt.Println(Go.Name)
	fmt.Println(Bd.Name)
	fmt.Printf("%+v\r\n", css)
	fmt.Printf("%+v\r\n", js)

	Go.ChangePrice(67.12) // (&Go).ChangePrice(67.12)
	Go.PrintClasses()     // (&Go).PrintClasses()

	fmt.Println(Go.Price)
}

/*
	func PrintClasses(c Course) {
		text := "Las clases son: "
		for _, class := range c.Classes {
			text += fmt.Sprintf("%s, ", class)
		}

		fmt.Println(text[:len(text)-2])
	}
*/
