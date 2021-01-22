package carrier

import (
	"errors"
	"fmt"

	"github.com/cesarcruzc/go-poo-from-scratch/ejercicio_ever/aircraft"
	"github.com/cesarcruzc/go-poo-from-scratch/ejercicio_ever/comms"
)

//AircraftCarrier struct
type AircraftCarrier struct {
	Aircrafts []aircraft.Aircraft
	comms     comms.Comms
}

//NewAircraftCarrier consructor
func NewAircraftCarrier(comms comms.Comms) *AircraftCarrier {
	return &AircraftCarrier{
		comms: comms,
	}
}

//Deploy method
func (ac *AircraftCarrier) Deploy(t string) (aircraft.Aircraft, error) {
	empty := aircraft.NewPlane("empty", "empty", "empty")

	for _, a := range ac.Aircrafts {
		if a.GetType() == t {
			err := ac.comms.Notify(a.GetType())
			if err != nil {
				fmt.Printf("Error Deploy %s", err)
				return empty, err
			}
			return a, nil
		}
	}

	return empty, errors.New("Aircraft Not found")
}

//AddAircraft method
func (ac *AircraftCarrier) AddAircraft(aircraft aircraft.Aircraft) {
	ac.Aircrafts = append(ac.Aircrafts, aircraft)
}

//ListAircrafts method
func (ac *AircraftCarrier) ListAircrafts() []aircraft.Aircraft {
	return ac.Aircrafts
}
