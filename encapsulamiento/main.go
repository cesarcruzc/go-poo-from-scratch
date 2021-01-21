package main

import (
	"fmt"

	"github.com/cesarcruzc/go-poo-from-scratch/encapsulamiento/course"
)

func main() {

	/*
		Constructor
	*/
	Go := course.New("Curso de GO", 0, false)

	/*
		Setters and Getters
	*/
	Go.SetName("POO con Go")
	Go.Name()

	Go.SetPrice(10.0)
	Go.Price()

	Go.SetIsFree(false)
	Go.IsFree()

	Go.SetUserIDs([]uint{12, 56, 89})
	Go.UserIDs()

	Go.SetClasses(map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	})
	Go.Classes() // (&Go).Classes()

	fmt.Printf("%+v\r\n", Go)

	/*
		Bd := course.Course{
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

		css := course.Course{
			Name:   "Curso de CSS",
			IsFree: false,
		}

		js := course.Course{}
		js.Name = "Curso JS"
		js.UserIDs = []uint{12, 56}

		fmt.Println(Go.Name)
		fmt.Println(Bd.Name)
		fmt.Printf("%+v\r\n", css)
		fmt.Printf("%+v\r\n", js)

		Go.ChangePrice(67.12) // (&Go).ChangePrice(67.12)
		Go.PrintClasses()     // (&Go).PrintClasses()

		fmt.Println(Go.Price)
	*/
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
