package main

import "fmt"

func main() {

	// Declaring a slice

	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice: %v\n" , slice)
	fmt.Printf("len of Slice: %v\n" , len(slice))

	// Slices work with pointers automatically

	newSlice := slice
	newSlice[1] = 20
	fmt.Printf("First slice: %v\n" , slice)
	fmt.Printf("Second slice: %v\n" , newSlice)

	// Getting sub-slices out of a slice

	u := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	w := u[:]
	x := u[:5]
	y := u[5:]
	z := u[3:7] // Intervals are like [a, b)
	fmt.Printf("Slices: %v, %v, %v, %v\n" , w, x, y, z)

	// Make a slice using "make" function

	made := make([]int, 3, 5) // type, len, cap
	fmt.Printf("Made slice: %v\n" , made)
	fmt.Printf("Made slice len: %v\n" , len(made))
	fmt.Printf("Made slice capacity: %v\n" , cap(made))

	// By adding a lot of elements to the slice, capacity gets increased too:

	made = append(made, 1)
	made = append(made, 1)
	made = append(made, 1)
	made = append(made, 1)
	made = append(made, 1)

	// Adding multiple elements to the slice

	made = append(made, 2, 3, 4, 5, 6, 7, 8, 9)
	made = append(made, []int{2, 3, 4, 5, 6, 7, 8, 9}...)

	fmt.Printf("New made slice: %v\n" , made)
	fmt.Printf("New made slice len: %v\n" , len(made))
	fmt.Printf("New made slice capacity: %v\n" , cap(made))

	// Removing an element

	temp := []int{1, 2, 3, 4, 5, 6}
	temp = append(temp[:2], temp[3:]...)
	fmt.Printf("slice without removed element: %v\n" , temp)

	// Weird situation (?)

	aa := []int{1, 2, 3, 4, 5}
	fmt.Println(aa)
	bb := append(aa[:2], aa[3:]...)
	fmt.Println(bb)
	fmt.Println(aa)
}
