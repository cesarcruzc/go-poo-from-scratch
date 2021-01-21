package main

import (
	"fmt"
	"strings"
)

type exampler interface {
	x()
}

// Son utiles para pasar cualquier tipo de dato
func wrapper(e interface{}) {
	fmt.Printf("value: %v, Type: %T\r\n", e, e)

	// Type assertions util to work with empty interfaces
	v, ok := e.(string)
	if ok {
		fmt.Println(strings.ToUpper(v))
	}

	switch v := e.(type) {
	case string:
		fmt.Printf("Parameter is string: %v", v)
	case int:
		fmt.Printf("Parameter is int: %v", v)
	case float64:
		fmt.Printf("Parameter is float: %v", v)
	default:
		fmt.Printf("Parameter is another: %v", v)
	}

}

func main() {
	// var e exampler
	// e.x()
	wrapper(1)
	wrapper(1.1)
	wrapper("hola")
}
