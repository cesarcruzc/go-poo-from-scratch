package carrier_test

// func TestCarrier(t *testing.T) {

// 	cases := []struct {
// 		name      string
// 		aircrafts []aircraft.Aircraft
// 		t         string // Type
// 		assert    func(result aircraft.Aircraft)
// 	}{
// 		{
// 			name: "successful deploy and attack",
// 			aircrafts: func() []aircraft.Aircraft {
// 				as := []aircraft.Aircraft{}
// 				as = append(as, aircraft.NewPlane("test f18", "Fighther", "Missil"))
// 				return as
// 			}(),
// 			t: "Fighther",
// 			assert: func(result aircraft.Aircraft) {
// 				// want := reflect.TypeOf(aircraft.Aircraft)
// 				// got := result

// 				want := "Launch Missil"
// 				got, _ := result.Attack()

// 				if got != want {
// 					t.Errorf("Error want: %s - got: %s", want, got)
// 				}

// 			},
// 		},
// 		{
// 			name: "failed attack",
// 			aircrafts: func() []aircraft.Aircraft {
// 				as := []aircraft.Aircraft{}
// 				as = append(as, aircraft.NewPlane("test f18", "Fighther", ""))
// 				return as
// 			}(),
// 			t: "Fighther",
// 			assert: func(result aircraft.Aircraft) {

// 				want := "Invalid weapon"
// 				_, err := result.Attack()

// 				if err.Error() != want {
// 					t.Errorf("Error want: %s - got: %s", want, err)
// 				}

// 			},
// 		},
// 		{
// 			name: "Successful bomber deploy and attack",
// 			aircrafts: func() []aircraft.Aircraft {
// 				bom := []aircraft.Aircraft{}
// 				bom = append(bom, aircraft.NewPlane("b2", "Bomber", "Bomb"))
// 				return bom
// 			}(),
// 			t: "Bomber",
// 			assert: func(plane aircraft.Aircraft) {
// 				want := "Launch Bomb"
// 				got, _ := plane.Attack()

// 				if got != want {
// 					t.Errorf("Error want:%s - got: %s", want, got)
// 				}
// 			},
// 		},
// 		{
// 			name: "Not found plane in carrier",
// 			aircrafts: func() []aircraft.Aircraft {
// 				bom := []aircraft.Aircraft{}
// 				bom = append(bom, aircraft.NewPlane("b2", "Bomber", "Bomb"))
// 				return bom
// 			}(),
// 			t: "Commercial",
// 			assert: func(result aircraft.Aircraft) {

// 				// t.Logf("%+v", result)

// 				want := "empty"
// 				got := result.GetType()

// 				if got != want {
// 					t.Errorf("Error want:%s - got: %s", want, got)
// 				}
// 			},
// 		},
// 	}

// 	for _, c := range cases {
// 		carrier := carrier.NewAircraftCarrier(c.aircrafts)
// 		plane := carrier.Deploy(c.t)
// 		c.assert(plane)
// 	}

// }
