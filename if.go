package main

import "fmt"

func main() {

	// If in go

	if x := 5 / 1; x > 3 {
		fmt.Printf("lol")
	}

	// Switch case in go (doesn't need "break" syntax)

	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("1, 5, 10")
	case 2, 3, 6:
		fmt.Println("2, 3, 6")
	default:
		fmt.Println("idk")
	}

	// Fallthrough (It's exactly the opposite of "break;" in Java or C++)

	switch i:= 2 + 3; {
	case i < 10:
		fmt.Println("less than 10")
		fallthrough
	case i < 20:
		fmt.Println("less than 20")
	default:
		fmt.Println("big")
	}

	// Switch over variable type

	var j interface{} = "1"

	switch j.(type) {
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("something")
	}
	
}