package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Creating a wait group to avoid using sleep in our main thread after starting a thread

var wg = sync.WaitGroup{}

// Variables for synchronization example

var waitGroup = sync.WaitGroup{}
var lock = sync.RWMutex{}
var counter = 0

func main() {

	// Starting a thread (we use sleep in order to be certain that the other thread is finished before main thread does)

	go printString("Hello, Go!")
	time.Sleep(100 * time.Millisecond)

	// Race condition

	var msg string = "Hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Bye"
	time.Sleep(100 * time.Millisecond)

	// Easy fix for race condition problem

	var newMsg string = "Hello"
	go func(msg string) {
		fmt.Println(msg)
	}(newMsg)
	newMsg = "Bye"
	time.Sleep(100 * time.Millisecond)

	// Using wait group

	var newerMsg string = "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(newerMsg)
	newerMsg = "Bye"
	wg.Wait()

	// Synchronizing

	for i := 0; i < 1000; i++ {
		waitGroup.Add(2)
		lock.RLock() // We should lock before calling these two functions to avoid deadlocks
		go printCounter()
		lock.Lock()
		go increaseCounter()
	}
	waitGroup.Wait()

	// Set max number of threads

	runtime.GOMAXPROCS(8)
	fmt.Printf("Threads: %v/n", runtime.GOMAXPROCS(-1))

}

func printString(str string)  {
	fmt.Println(str)
}

func printCounter()  {
	fmt.Printf("Counter: %v\n", counter)
	lock.RUnlock()

	waitGroup.Done()
}

func increaseCounter()  {
	counter++
	lock.Unlock()

	waitGroup.Done()
}