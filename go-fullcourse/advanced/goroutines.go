package advanced

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{} // Create a new WaitGroup to make GoRoutines.

func GoRoutines_Concept(){
	var msg = "Hello"

	wg.Add(1) // Add a new routine to the main function

	go func(msg string){ // Spawn a goroutine (a new green thread)
		fmt.Println(msg)
		wg.Done() // Acknowledge the function ended.
	}(msg)

	msg = "Goodbye"

	wg.Wait() // Wait the end of all go routines to finish the main function.
}

// #####################################################

var wg2 = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func GoRoutines_SynchronizingRoutines(){
	for i := 0; i < 10; i++ {
		wg2.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}

	wg2.Wait()
}

func sayHello(){
	fmt.Printf("Hello %v\n", counter)
	m.RUnlock()
	wg2.Done()
}

func increment(){
	counter++
	m.Unlock()
	wg2.Done()
}

func GoRoutines_MAXPROCS(){
	runtime.GOMAXPROCS(2) // set the number of threads for this program
	fmt.Printf("Threads: %v", runtime.GOMAXPROCS(-1))
}