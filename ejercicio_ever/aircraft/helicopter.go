package aircraft

import (
	"errors"
	"fmt"
)

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
