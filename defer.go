package main

import "fmt"

func main() {

	// defer runs at the end of function, before returning, from down to top

	// Example 1

	fmt.Println("0")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("4")

	// Example 2

	str := "hi"
	defer fmt.Println(str) // prints "hi"
	str = "bye"
}
