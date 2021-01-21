package main

import (
	"errors"
	"fmt"
)

// Aircraft interface
type Aircraft interface {
	GetType() string
	Flight() error
	Attack() (string, error)
}

type aircraft struct {
	Name   string
	Type   string
	Weapon string
}

// NewAircraft cosntructor
func NewAircraft(name, t, weapon string) Aircraft {
	return &aircraft{Name: name, Type: t, Weapon: weapon}
}

func (a aircraft) GetType() string {
	return a.Type
}

func (a aircraft) Flight() error {
	return errors.New("Error Flight")
}

func (a aircraft) Attack() (string, error) {
	if a.Weapon != "Missil" {
		return "", errors.New("Invalid weapon")
	}

	return fmt.Sprintf("Launch %s\r\n", a.Weapon), nil
}

func (a aircraft) Landing() error {
	return errors.New("Error Landing")
}

func (a aircraft) checkEngine() error {
	return errors.New("Error CheckEngine")
}

//AircraftCarrier struct
type AircraftCarrier struct {
	Aircrafts []Aircraft
}

//NewAircraftCarrier consructor
func NewAircraftCarrier(as []Aircraft) *AircraftCarrier {
	return &AircraftCarrier{
		Aircrafts: as,
	}
}

//Deploy method
func (ac *AircraftCarrier) Deploy(t string) Aircraft {
	for _, a := range ac.Aircrafts {
		if a.GetType() == t {
			fmt.Printf("%+v", a)
			return a
		}
	}

	return aircraft{}
}

type helicopter struct {
	Name   string
	Type   string
	Weapon string
}

//NewHelicopter constuctor
func NewHelicopter(name, t, weapon string) Aircraft {
	return &helicopter{Name: name, Type: t, Weapon: weapon}
}

func (h helicopter) GetType() string {
	return h.Type
}

func (h helicopter) Flight() error {
	return errors.New("Flight Helicopter Error")
}

func (h helicopter) Attack() (string, error) {
	return fmt.Sprintf("Using my %s ", h.Weapon), nil
}

func main() {

	f17 := NewAircraft("f17", "Fighter", "Missil")
	fmt.Printf("%+v\r\n", f17.GetType())
	// fmt.Printf("%+v\r\n", f17.Landing())
	// fmt.Printf("%+v\r\n", f17.checkEngine())

	f16 := aircraft{"f16", "Fighter", "Missil"}
	fmt.Printf("%+v\r\n", f16.GetType())
	fmt.Printf("%+v\r\n", f16.Landing())

	b52 := NewAircraft("b52", "Bomber", "Bomb")

	fmt.Println("=====================")

	apache := NewHelicopter("Apache", "Helicopter", "Machine Gun")

	aircrafts := []Aircraft{}
	aircrafts = append(aircrafts, f16)
	aircrafts = append(aircrafts, f17)
	aircrafts = append(aircrafts, b52)
	aircrafts = append(aircrafts, apache)

	carrier := NewAircraftCarrier(aircrafts)
	fmt.Printf("%+v\r\n", carrier)

	fmt.Println("=====================")

	a1 := carrier.Deploy("Fighter")

	f, err := a1.Attack()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(f)

	// fmt.Printf("%s\r\n", )

	a2 := carrier.Deploy("Bomber")

	f2, err := a2.Attack()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(f2)

	fmt.Println("=====================")
	ap := carrier.Deploy("Helicopter")
	f3, err := ap.Attack()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(f3)

}
