/*	Channels in go are used to receive and send data to and from go routines */

package main

import ("fmt"
"time")

// func sendData(mychan chan string) {
// 	/this channel can send and recieve data both ways if we want to make it one way then we have
// 	/ to specify the direction of the channel
// 	/ chan string is the direction of the channel
// 	/ chan<- string is the direction of the channel that can only send data
// 	/ <-chan string is the direction of the channel that can only receive data
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
// func sum(resultChan chan int, num1, num2 int) {
// 	resultChan <- num1 + num2
// }
//go routine synchronization
//for synchronization instead of wait groups
// func task(done chan bool){
	//defer is a function that runs in the last whether the func cleanly run or give an error
// 	defer func(){
// 		done <-true
// 	}()
// 	fmt.Println("Processing done")
// }

func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}
      
	//normal channels are blocking but buffer channles allow us to send a an amount of data
	//without blocking
	// emailChan := make(chan string,100)//100 buffer channle like we can send 100 channels without blocking
	// done := make(chan bool)
	// go emailchan(emailChan,done)

	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }
 
	// fmt.Println("done")

	//for a channel we have to use the in built close function
	// close(emailChan)
	// <-done








	// numChan := make(chan int)
	// done := make(chan bool)

	// go sum(numChan, 5, 5)

	// res := <-numChan //blocking call

	// fmt.Println(res)

	// go task(done)
    // now we are using it as a wait group use channels for only one go if multiple use wait group for clean syntax
	// <- done

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
