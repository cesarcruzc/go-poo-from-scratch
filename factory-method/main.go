package main

import "fmt"

//PaymentMethod interface
type PaymentMethod interface {
	Pay()
}

//Paypal struct
type Paypal struct{}

//Pay method to pay with PayPal
func (p Paypal) Pay() {
	fmt.Println("Payed by Paypal")
}

//Cash struct
type Cash struct{}

//Pay method to pay with Cash
func (c Cash) Pay() {
	fmt.Println("Payed by Cash")
}

//CreditCard struct
type CreditCard struct{}

//Pay method to pay with Credit Card
func (c CreditCard) Pay() {
	fmt.Println("Payed by credit card")
}

//Factory function to return the interface according to the selected option
func Factory(method uint) PaymentMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default:
		return nil
	}
}

func main() {

	var method uint
	fmt.Println("Select a payment method")
	fmt.Println("\t 1: Paypal 2:Cash, 3:Credit card")

	_, err := fmt.Scanln(&method)
	if err != nil {
		panic("You have to select a valid payment method")
	}

	if method > 3 {
		panic("You have to select a valid payment method")
	}

	payMethod := Factory(method)
	payMethod.Pay()
}
