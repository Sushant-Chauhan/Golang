package main

import "fmt"

//-------------------------------

type payment struct{}

func (p payment) makePayment(amount float32) {
	razorpayPaymentGtw := razorpay{}
	razorpayPaymentGtw.pay(amount)

	stripePaymentGtw := stripe{}
	stripePaymentGtw.pay(amount)

}

//-------------------------------

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	//logic to make payment (call razorpay apis)
	fmt.Println("making payment using razorpay - ", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe - ", amount)
}

//-------------------------------

func main() {
	newPayment1 := payment{}
	newPayment1.makePayment(500)

	newPayment2 := payment{}
	newPayment2.makePayment(800)

}
