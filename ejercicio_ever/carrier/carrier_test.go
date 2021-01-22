package carrier_test

import (
	"errors"
	"testing"

	"github.com/cesarcruzc/go-poo-from-scratch/ejercicio_ever/aircraft"
	"github.com/cesarcruzc/go-poo-from-scratch/ejercicio_ever/carrier"
	"github.com/cesarcruzc/go-poo-from-scratch/ejercicio_ever/comms"
)

type commsMock struct {
	// NotifyMock func(message string) error
}

// func New() {}

func (c commsMock) Notify(message string) error {
	// return c.NotifyMock(message)
	return nil
}

type commsMock2 struct {
	// NotifyMock func(message string) error
}

// func New() {}

func (c commsMock2) Notify(message string) error {
	return errors.New("Error mock2")
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
				// want := reflect.TypeOf(aircraft.Aircraft)
				// got := result

				want := "Launch Missil"
				got, _ := result.Attack()

				if got != want {
					t.Errorf("Error want: %s - got: %s", want, got)
				}

			},
			comms: commsMock{},
			// comms: commsMock{
			// 	NotifyMock: func(message string) error {
			// 		return nil
			// 	},
			// },
		},
		{
			name:     "failed deploy",
			aircraft: aircraft.NewPlane("test f18", "Fighther", ""),
			t:        "Fighther",
			assert: func(result aircraft.Aircraft, err error) {
				if err != nil {
					t.Errorf("Error want: %s - got: %v", err, nil)
				}

			},
			comms: commsMock2{},
		},

		// {
		// 	name:     "failed attack",
		// 	aircraft: aircraft.NewPlane("test f18", "Fighther", ""),
		// 	t:        "Fighther",
		// 	assert: func(result aircraft.Aircraft, err error) {

		// 		want := "Invalid weapon"
		// 		_, e := result.Attack()

		// 		if e.Error() != want {
		// 			t.Errorf("Error want: %s - got: %s", want, e)
		// 		}

		// 	},
		// 	comms: commsMock2{},
		// },

		// {
		// 	name: "Successful bomber deploy and attack",
		// 	aircrafts: func() []aircraft.Aircraft {
		// 		bom := []aircraft.Aircraft{}
		// 		bom = append(bom, aircraft.NewPlane("b2", "Bomber", "Bomb"))
		// 		return bom
		// 	}(),
		// 	t: "Bomber",
		// 	assert: func(plane aircraft.Aircraft) {
		// 		want := "Launch Bomb"
		// 		got, _ := plane.Attack()

		// 		if got != want {
		// 			t.Errorf("Error want:%s - got: %s", want, got)
		// 		}
		// 	},
		// 	comms: commsMock{},
		// },
		// {
		// 	name: "Not found plane in carrier",
		// 	aircrafts: func() []aircraft.Aircraft {
		// 		bom := []aircraft.Aircraft{}
		// 		bom = append(bom, aircraft.NewPlane("b2", "Bomber", "Bomb"))
		// 		return bom
		// 	}(),
		// 	t: "Commercial",
		// 	assert: func(result aircraft.Aircraft) {

		// 		// t.Logf("%+v", result)

		// 		want := "empty"
		// 		got := result.GetType()

		// 		if got != want {
		// 			t.Errorf("Error want:%s - got: %s", want, got)
		// 		}
		// 	},
		// 	comms: commsMock{},
		// },
	}

	for _, c := range cases {
		carrier := carrier.NewAircraftCarrier(c.comms)
		carrier.AddAircraft(c.aircraft)
		plane, err := carrier.Deploy(c.t)
		c.assert(plane, err)
	}

}
