package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

// order struct
type order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time //nano second precision
	customer            //STRUCT EMBEDED IN CODE - composition
}

// so Important - struct, embeded, composition

func newOrder(id string, amount float32, status string) *order {
	//initial setup goes here...
	myOrder := order{
		id:     id,
		amount: float64(amount),
		status: status,
	}
	return &myOrder // return address of struct
}

// receiver type
func (o *order) changedStatus(status string) {
	o.status = status
}

func (o *order) getAmount() float32 {
	return float32(o.amount)
}

func main() {

	// newCustomer := customer{
	// 	name:  "john",
	// 	phone: "1234567890",
	// }
	myorder4 := order{
		id:     "1",
		amount: 30,
		status: "received",
		// customer: newCustomer,
		customer: customer{ //inline assign
			name:  "john",
			phone: "1234567890",
		},
	}
	fmt.Println("myorder4 = ", myorder4)
	myorder4.name = "ROBIN"
	fmt.Println("myorder4 = ", myorder4)

	//Inline struct
	language := struct {
		name   string
		isGood bool
	}{"golang", true}
	fmt.Println(language)

	myorder3 := newOrder("3", 30.50, "REceived")
	fmt.Println("myorder3 = ", myorder3)
	fmt.Println("myorder3 amount = ", myorder3.amount) //no need to dereference, struct will do it itself

	myorder1 := order{
		id:     "1",
		amount: 50.534,
		status: "received",
	}
	myorder1.createdAt = time.Now()

	myorder2 := order{
		id:     "2",
		amount: 100.534,
		status: "paid",
	}
	myorder2.changedStatus("Confirmed")

	fmt.Println("myorder1 = ", myorder1)
	fmt.Println("myorder1 amount = ", myorder1.getAmount())
	fmt.Println("myorder2 = ", myorder2)
}
