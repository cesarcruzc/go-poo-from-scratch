package main

import (
	"fmt"

	"github.com/cesarcruzc/go-poo-from-scratch/ejercicio_ever/aircraft"
	"github.com/cesarcruzc/go-poo-from-scratch/ejercicio_ever/carrier"
	"github.com/cesarcruzc/go-poo-from-scratch/ejercicio_ever/comms"
)

func main() {

	comms := comms.New()
	carrier := carrier.NewAircraftCarrier(comms)

	f16 := aircraft.NewPlane("f16", "Fighther", "Missil")
	carrier.AddAircraft(f16)

	// fmt.Printf("%+v", carrier.ListAircrafts())

	a1 := carrier.Deploy("Fighther")

	result, err := a1.Attack()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	fmt.Printf("%+v", result)

}
