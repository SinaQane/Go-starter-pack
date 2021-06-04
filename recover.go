package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("hi")
	panicHandler() // Using this, app recovers after panic
	fmt.Println("bye")
}

func panicHandler()  {
	fmt.Print("starting to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)

			// We can use the following if/else statement instead of the line above:

			/*
			if WE CAN'T HANDLE THE ERROR {
				panic(err)
			} else {
				log.Println(err) // If we can handle the panic and we want app to keep running...
			}
			 */

		}
	}()
	panic("panicking")
	fmt.Println("end of panic")
}
