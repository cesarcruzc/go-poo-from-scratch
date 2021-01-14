package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func main() {
	Go := Course{
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
}
