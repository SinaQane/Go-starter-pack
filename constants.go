package main

import (
	"fmt"
)

// Declaring a constant in package scope

const a int16 = 0

// iota in const block

const (
	x = iota
	y = iota
	z = iota
)

// iota gets reset in another block

const (
	m = iota
	n
	p
)

// Ignore the zero-th index

const (
	_ = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	YB
)

// Shift iota 5 indices and ignore the zero-th index as well

const (
	_ = iota + 5
	first
	second
)

func main ()  {

	// Declaring a constant in main function scope

	const myConst int = 1
	fmt.Println(myConst)

	/*
		const pi float64 = math.Sin(1.57) // Gets an error, since it is the result of a function and not a true constant
	 */

	// Re-declaring a constant that has already been declared in package scope

	const a int32 = 10
	fmt.Println(a)

	// Printing constants that were created by iota

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(p)

	fmt.Println(first)
}
