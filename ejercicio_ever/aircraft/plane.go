package aircraft

import (
	"errors"
	"fmt"
)

type plane struct {
	Name   string
	Type   string
	Weapon string
}

//NewPlane constructor
func NewPlane(name, t, weapon string) Aircraft {
	return &plane{Name: name, Type: t, Weapon: weapon}
}

func (a plane) GetType() string {
	return a.Type
}

func (a plane) Flight() error {
	return errors.New("Error Flight")
}

func (a plane) Attack() (string, error) {
	if a.Weapon == "" {
		return "", errors.New("Invalid weapon")
	}

	return fmt.Sprintf("Launch %s", a.Weapon), nil
}

func (a plane) Landing() error {
	return errors.New("Error Landing")
}

func (a plane) checkEngine() error {
	return errors.New("Error CheckEngine")
}
