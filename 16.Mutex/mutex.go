//mutex are used to decrease the risk of overwriting of resource by another go routine
// lets say we have multiple go routine modifiying a single resource so there is a risk that
// one go routine may override the changes of one etc so we use mutex for that purpose
// it comes from the package sync

package main

import (
	"fmt"
	"sync"
)


type post struct{
	view int
	mu sync.Mutex // you can declare it anywhere global 
}

func (p*post) inc(wg * sync.WaitGroup){
	// defer wg.Done()
    // defer p.mu.Unlock()
	//single deffer are lifo means first defer will be the last to execute
	//func main() {
// 	defer fmt.Println("1")
// 	defer fmt.Println("2")
// 	defer fmt.Println("3")
// }
// Output:
// 3
// 2
// 1
	// Inside a single deferred closure, the statements run in the normal top-to-bottom order you wrote them — 
	// that part looks FIFO-ish because it's just... normal sequential code execution inside one function body:
// 	defer func() {
// 	fmt.Println("A") // runs first
// 	fmt.Println("B") // runs second
// }()
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock()
	p.view++ // without mutex all go routines are changing the value of the 
	// same resource we can get less value than 100 so mutex is necessary here
	// you can unlock the mutex here but lets take example if there comes a error so the program cant reach
	// to the unlock so it is better to unlock it in defer becaue defer always run 
}

func main(){
	var wg sync.WaitGroup
	post1:= post{view:0}

	for i:=0; i<=100; i++{
		wg.Add(1)
        go post1.inc(&wg)
	}

	wg.Wait()
	fmt.Println("The total view count:",100)
}