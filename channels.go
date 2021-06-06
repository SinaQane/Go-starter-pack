package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup for our goroutines

var wait = sync.WaitGroup{}

// Channels for logging example

var logChannel = make(chan logEntry, 100)
var doneChannel = make(chan struct{}) // Signal only channel (like "make(chan bool)" but better)

func main() {

	// Channels are designed to sync data transmission between multiple goroutines

	channel := make(chan int) // chan = channel, int = datatype that's going to float into the channel
	wait.Add(2)
	go func() {
		i := <- channel
		fmt.Println(i)
		wait.Done()
	}()
	go func() {
		num := 69
		channel <- num
		num = 420
		wait.Done()
	}()
	wait.Wait()

	// By pushing and not using data into the channels you can cause a dead-lock

	wait.Add(6)
	go func() {
		i := <- channel
		fmt.Println(i)
		wait.Done()
	}()
	for i := 0; i < 5; i++ {
		go func() {
			num := 69
			channel <- num // Pauses the execution of goroutine until there's space available in channel
			wait.Done()
		}()
	}
	wait.Wait()

	// More complicated example

	wait.Add(2)
	go func() {
		fmt.Println(<- channel)
		channel <- 420
		wait.Done()
	}()
	go func() {
		channel <- 69
		fmt.Println(<- channel)
		wait.Done()
	}()
	wait.Wait()

	// Modification (Send & Receive only channels)

	wait.Add(2)
	go func(channel <-chan int) { // Receive only channel
		fmt.Println(<- channel)
		wait.Done()
	}(channel)
	go func(channel chan<- int) { // Send only channel
		channel <- 69
		wait.Done()
	}(channel)
	wait.Wait()

	// Increasing the storage size of our channel

	newChannel := make(chan int, 10) // 10 is the storage size of newChannel (it can store 10 integers)

	wait.Add(2)
	go func(channel <-chan int) { // Receive only channel
		for i := range newChannel {
			fmt.Println(i)
		}

		/* Another approach:

		for {
			if i, ok := <- newChannel; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		*/

		wait.Done()
	}(newChannel)
	go func(channel chan<- int) { // Send only channel
		channel <- 69
		channel <- 6969
		channel <- 696969
		channel <- 69696969
		close(newChannel) // Closing channel to avoid deadlock in receiver function (You can't re-open channel and send data to it)
		wait.Done()
	}(newChannel)
	wait.Wait()

	// Logging with channels example

	/* We could have done this instead of using doneChannel (but not recommended for every situation):

	defer func() {
		close(logChannel)
	}
	*/
	go logger()

	logChannel <- logEntry{time.Now(), "INFO", "Starting..."}
	logChannel <- logEntry{time.Now(), "INFO", "Stopping..."}

	time.Sleep(100 * time.Millisecond)

	doneChannel <- struct{}{}
}

type logEntry struct {
	time        time.Time
	sensitivity string
	message     string
}

func logger()  {

	/* While not using doneChannel:

	for entry := range logChannel {
		fmt.Printf("%v - [%v] %v\n", entry.time, entry.sensitivity, entry.message)
	}
	*/

	for {
		select {
		case entry := <- logChannel:
			fmt.Printf("%v - [%v] %v\n", entry.time, entry.sensitivity, entry.message)
		case <- doneChannel:
			break
		// default:
		}
	}
}
