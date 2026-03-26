package main

import "fmt"

type Payment interface {
	pay(amount float64)
}

type Creditcard struct{}
type Upi struct{}
type Paypal struct{}

func (c Creditcard) pay(amount float64) {
	fmt.Println("payment done using creditcard ", amount)
}

func (u Upi) pay(amount float64) {
	fmt.Println("Payment done using UPI ", amount)
}

func (p Paypal) pay(amount float64) {
	fmt.Println("Payment done using Paypal ", amount)
}

func main() {
	var c Payment
	var u Payment
	var p Payment

	c = Creditcard{}
	u = Upi{}
	p = Paypal{}

	c.pay(1000)
	u.pay(2000)
	p.pay(2000)

}
