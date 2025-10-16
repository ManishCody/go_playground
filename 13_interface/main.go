package main

import "fmt"

type Paymentor interface {
	pay(amount float64)
}

type payment struct {
	getway Paymentor
}

func (o payment) makePayment(amount float64) {
	// razorpayPayments := razorpay{}
	// paypalPayments := paypal{}
	// razorpayPayments.pay(amount)
	o.getway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float64) {
	fmt.Println("Payment of", amount, "made using Razorpay")
}

type paypal struct{}

func (p paypal) pay(amount float64) {
	fmt.Println("Payment of", amount, "made using Paypal")
}

func main() {
	razorpayPayment := razorpay{}
	paypalPayment := paypal{}
	newPayment := payment{
		getway: paypalPayment,
	}
	newPayment2 := payment{
		getway: razorpayPayment,
	}
	newPayment2.makePayment(3000)
	newPayment.makePayment(5000)
}
