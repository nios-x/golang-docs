package main

import "fmt"

type paymenter interface {
	pay(x float32)
}

type razorpay struct{}
type gpay struct{}

func (e razorpay) pay(amt float32) {
	fmt.Println("Paid", amt, "throungh Razorpay")
}
func (e gpay) pay(amt float32) {
	fmt.Println("Paid", amt, "throungh Gpay")
}

type payment struct {
	paymentclass paymenter
}

func (e payment) pay(amount float32) {
	e.paymentclass.pay(amount)
}
func main() {
	razorpayobj := razorpay{}
	gpayobj := gpay{}
	gateway := payment{
		paymentclass: razorpayobj,
	}
	gateway.pay(77)
	gateway.paymentclass = gpayobj
	gateway.pay(347)
}
