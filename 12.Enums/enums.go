package main

import "fmt"

//in go we dont have enums like we do in cpp java etc
// instead we can use the const as a enums

type order int

const (
	Recieved   order = iota // we can use it with int type it value is incremented automatically
	Confirmed               //1
	Prepared                //2
	Deleieverd              //3
)

// we can also assign this an sting we just have to change
// type order string
// const (
// 		Recieved order = "recieved"
// 		Confirmed      = "confirmed"
//    	.......)

func orderStatus(status order) {
	fmt.Println("the order status is; ", status)
}

func main() {
	orderStatus(Prepared)
}
