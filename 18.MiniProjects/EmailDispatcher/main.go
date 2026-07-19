package main

import "sync"

// entry point of our application

type Recipent struct {
	Name  string
	Email string
}

func main() {

	recipentChan := make(chan Recipent)

	go loadRecipent("contacts.csv", recipentChan)

	var wg sync.WaitGroup

	workerCount := 5
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipentChan, &wg)
	}

	

	wg.Wait()
}
