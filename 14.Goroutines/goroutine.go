package main

import (
	"fmt"
	"sync"
)

func print(i int, wg *sync.WaitGroup) {
	//3. after executing the parallel task delete it or make it done the task from the waitgroup
	// Decrements the counter by 1 (called when a goroutine finishes)
	// Defer defer postpones a function call until the surrounding function is about to return —
	// commonly used for cleanup tasks, guaranteeing they run regardless of how the function exits
	// (normal return, early return, or panic).
	defer wg.Done() //defer means it execautes after the function
	fmt.Println(i)
}

func main() {
	// for i := 0; i <= 10; i++ {
	// 	go print(i) // the normal without go it will block the code untill print(0) didnt run full
	// 	// but with the use of go routines the all functions will run in parallel on
	// 	// different light weight thread and doesnt print in order
	// }
	// // now there is issue with we are passing time and we dont when our function or parallel task will
	// // stop so we can use the waitgroup for that
	// time.Sleep(time.Second * 1) // delay it for one second because main is also running in one thread
	// // when we run all above task in parallel there is no task that will block the main and when main
	// // comes below there will no code and program end so we delay to see the output of the go print

	//1 first create the waitgroup from sync package
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		//2 add the task in the wait group
		// Increments the internal counter by n (how many goroutines to wait for)
		wg.Add(1)
		go print(i, &wg)
	}

	//4.it's used to block/pause execution until a group of goroutines finish running
	wg.Wait()
}
