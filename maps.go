package main

import (
	"fmt"
)

func main() {

	// Creating a map

	m := make(map[string]int)
	m = map[string]int{
		"hi" : 100,
		"bye" : 99,
	}

	// Key and value in map can't be every object, for example slices can't be used in maps

	/*
		tempMap := make(map[[]int]int) // Gets an error
	 */

	// Items in map will not be printed in order later while iterating or printing...

	fmt.Println(m)
	fmt.Println(m["hi"])

	m["lol"] = 1000
	delete(m, "hi")

	fmt.Println(m)

	fmt.Println(m["yo"]) // This will print out 0.

	// To find out if the key doesn't exist or it's value is 0 do this:

	pop, ok := m["yo"]
	fmt.Println(pop, ok)

	// Maps have built-in pointer references like slices

	map1 := map[string]int{
		"a" : 1,
		"b" : 2,
		"c" : 3,
		"d" : 4,
	}
	map2 := map1
	delete(map2, "b")
	fmt.Println(map1)
	fmt.Println(map2)
}