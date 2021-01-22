package carrier

import (
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
func (ac *AircraftCarrier) Deploy(t string) aircraft.Aircraft {
	for _, a := range ac.Aircrafts {
		if a.GetType() == t {
			ac.comms.Notify(a.GetType())
			return a
		}
	}

	return aircraft.NewPlane("empty", "empty", "empty")
}

//AddAircraft method
func (ac *AircraftCarrier) AddAircraft(aircraft aircraft.Aircraft) {
	ac.Aircrafts = append(ac.Aircrafts, aircraft)
}

//ListAircrafts method
func (ac *AircraftCarrier) ListAircrafts() []aircraft.Aircraft {
	return ac.Aircrafts
}
