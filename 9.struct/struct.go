package main

import (
	"fmt"
	"time"
)

type order struct {
		order_id int
		name string
		createdAt time.Time
	}


func main(){

	myOrder:= order{
		order_id: 1,
		name: "Hoodie"
	}

	fmt.Println(myOrder)
}