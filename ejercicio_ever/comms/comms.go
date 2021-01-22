package comms

import (
	"errors"
	"fmt"
)

//Comms interface
type Comms interface {
	Notify(message string) error
}

type comms struct{}

func New() Comms {
	return comms{}
}

func (c comms) Notify(message string) error {
	if message == "" {
		return errors.New("Invalid message")
	}

	fmt.Printf("A %s has been deployed successfully", message)

	return nil

}
