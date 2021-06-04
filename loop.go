package main

import "fmt"

func main() {

	// While loop in go (There's no "while" or "do/while" syntax)

	var i int = 0
	for i < 5 { // or "; i < 5 ;"
		fmt.Printf("1st loop: %v\n", i)
		i++
	}

	// Infinite loop (Again, there's no "while" for infinite loops)

	var j int = 0
	for {
		fmt.Printf("2nd loop: %v\n", j)
		j++
		if j > 10 {
			break
		}
	}

	// Continue

	for k := 0; k < 10; k++ {
		if k % 2 == 0 {
			continue
		}
		fmt.Printf("3rd loop: %v\n", k)
	}

	// Label loops for breaking out of them

	loop1:
		for a := 0; a < 3; a++ {
			// loop2:
			for b := 0; b < 3; b++ {
				fmt.Println(a * b)
				if a*b > 2 {
					break loop1
					// break loop2
				}
			}
		}

	// Iterate over a slice with for ("for each" syntax)

	slc := []int{1, 2, 3, 4, 5}

	for key, value := range slc {
		fmt.Println(key, value)
	}

	// Channels are used for concurrency

}