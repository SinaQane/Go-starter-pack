package main

import (
	"fmt"
)

func main ()  {

	// Creating arrays and working with them

	grades := [...]int{1, 10, 100}
	fmt.Printf("%v, %v, %v\n" , grades[0], grades[1], grades[2])

	var students [3]string
	fmt.Printf("Students: %v\n" , students)
	students[0] = "Sina"
	students[1] = "Sina2"
	students[2] = "Sina3"
	fmt.Printf("Students: %v\n" , students)
	fmt.Printf("Student #1: %v\n" , students[0])
	fmt.Printf("no of Students: %v\n" , len(students))

	a := [...]int{1, 2, 3}

	// Changing a value in array without pointers

	b := a
	b[1] = 20
	fmt.Printf("First array: %v\n" , a)
	fmt.Printf("Second array: %v\n" , b)

	// Changing a value in array using pointers

	c := &a
	c[1] = 40
	fmt.Printf("First array: %v\n" , a)
	fmt.Printf("Second array: %v\n" , c)
}
