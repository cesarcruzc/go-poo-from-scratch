package carrier_test

import (
	"errors"
	"testing"

	"github.com/cesarcruzc/go-poo-from-scratch/ejercicio_ever/aircraft"
	"github.com/cesarcruzc/go-poo-from-scratch/ejercicio_ever/carrier"
	"github.com/cesarcruzc/go-poo-from-scratch/ejercicio_ever/comms"
)

type commsMock struct {
	NotifyMock func(message string) error
}

func (c commsMock) Notify(message string) error {
	return c.NotifyMock(message)
}

func TestCarrier(t *testing.T) {

	cases := []struct {
		name     string
		aircraft aircraft.Aircraft
		t        string // Type
		assert   func(result aircraft.Aircraft, err error)
		comms    comms.Comms
	}{
		{
			name:     "successful deploy and attack",
			aircraft: aircraft.NewPlane("test f18", "Fighther", "Missil"),
			t:        "Fighther",
			assert: func(result aircraft.Aircraft, err error) {
				want := "Launch Missil"

				got, _ := result.Attack()
				if got != want {
					t.Errorf("Error want: %s - got: %s", want, got)
				}

			},
			comms: commsMock{
				NotifyMock: func(message string) error {
					return nil
				},
			},
		},
		{
			name:     "failed deploy",
			aircraft: aircraft.NewPlane("test f18", "Fighther", ""),
			t:        "Fighther",
			assert: func(result aircraft.Aircraft, err error) {
				if err == nil {
					t.Errorf("Error want: %s - got: %v", err, nil)
				}

			},
			comms: commsMock{
				NotifyMock: func(message string) error {
					return errors.New("Error failed send message")
				},
			},
		},
		{
			name:     "failed attack",
			aircraft: aircraft.NewPlane("test f18", "Fighther", ""),
			t:        "Fighther",
			assert: func(result aircraft.Aircraft, err error) {
				want := "Invalid weapon"
				_, e := result.Attack()
				if e.Error() != want {
					t.Errorf("Error want: %s - got: %s", want, e)
				}
			},
			comms: commsMock{
				NotifyMock: func(message string) error {
					return nil
				},
			},
		},
		{
			name:     "Successful bomber deploy and attack",
			aircraft: aircraft.NewPlane("b2", "Bomber", "Bomb"),
			t:        "Bomber",
			assert: func(plane aircraft.Aircraft, err error) {
				want := "Launch Bomb"
				got, _ := plane.Attack()

				if got != want {
					t.Errorf("Error want:%s - got: %s", want, got)
				}
			},
			comms: commsMock{
				NotifyMock: func(message string) error {
					return nil
				},
			},
		},
		{
			name:     "Not found plane in carrier",
			aircraft: aircraft.NewPlane("AirBus", "Commercial", ""),
			t:        "Bomber",
			assert: func(result aircraft.Aircraft, err error) {

				if err == nil {
					t.Errorf("Error %+v", err)
				}
			},
			comms: commsMock{
				NotifyMock: func(message string) error {
					return nil
				},
			},
		},
	}

	for _, c := range cases {
		carrier := carrier.NewAircraftCarrier(c.comms)
		carrier.AddAircraft(c.aircraft)
		plane, err := carrier.Deploy(c.t)

		c.assert(plane, err)
	}

}
