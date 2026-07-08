package main

import (
	"fmt"
	"time"
)

type order struct {
	order_id  int
	name      string
	createdAt time.Time
}

// in go there are no constuctors so we use a trick like this

func newOrder(order_id int, name string) *order {
	myOrder := order{
		order_id:  order_id,
		name:      name,
		createdAt: time.Now(),
	}
	return &myOrder
}

// Reciever function

func (o *order) changedName(newName string) {
	//whenever we have to change some value we have to use the dereference operator in the reciever function
	o.name = newName // *o.name we dont have to dereference here struct will do for use
}

func (o order) getOrderId() int {
	//as we are not changing anything so no dont have to use the dereference operator
	return o.order_id
}

func main() {

	// in some scenario we have to make only one instance of the struct
	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

	// if you dont set any field default value will be automatically set
	myOrder := order{
		order_id: 1,
		name:     "Hoodie",
	}
	// we can create the multiple instance of the struct order
	// myOrder2 := order{
	// 	order_id:  2,
	// 	name:      "shirt",
	// 	createdAt: time.Now(),
	// }

	//constructor equivalent function
	myOrder3 := newOrder(3, "hoodie")
	myOrder.changedName("jeans")
	fmt.Println(myOrder.getOrderId())
	// myOrder.createdAt = time.Now()
	fmt.Println(myOrder)
	fmt.Println(myOrder3)
}
