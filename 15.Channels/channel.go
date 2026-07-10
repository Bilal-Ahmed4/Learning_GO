/*	Channels in go are used to receive and send data to and from go routines */

package main

import "fmt"

// func sendData(mychan chan string) {
// 	//this channel can send and recieve data both ways if we want to make it one way then we have
// 	// to specify the direction of the channel
// 	// chan string is the direction of the channel
// 	// chan<- string is the direction of the channel that can only send data
// 	// <-chan string is the direction of the channel that can only receive data
// 	mychan <- "go lang " //Now we are sending data to the channel
// }

//send
// func processNum(numChan chan int) {
// 	for num := range numChan { //range will automatically receive data from the channel
// 		fmt.Println("channels processed", num)
// 		time.Sleep(time.Second)
// 	}
// }

// receive
func sum(resultChan chan int, num1, num2 int) {
	resultChan <- num1 + num2
}

func main() {
	numChan := make(chan int)

	go sum(numChan, 5, 5)

	res := <-numChan //blocking call

	fmt.Println(res)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	//time.Sleep(time.Second)

	// mychan <- "go lang " //Now we are sending data to the channel // this is a blocking call block till the other go routine is ready to receive
	// go sendData(mychan)
	// data := <-mychan //Now we are receiving data from the channel
	// fmt.Println(data)

}
