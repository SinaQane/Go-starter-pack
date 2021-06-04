package main

import "fmt"

func main() {

	// Change a variable without pointers

	a := 20
	b := a
	fmt.Println(a, b)
	a = 30
	fmt.Println(a, b)

	// Change a variable with pointers

	var x int = 20
	var y *int = &x
	fmt.Println(x, y)
	x = 30
	fmt.Println(x, *y)

	// Example 1

	var obj *newStruct
	obj = new(newStruct)
	(*obj).foo = 100
	fmt.Println((*obj).foo)

		// Last two lines in this example can be written like this:

	obj.foo = 100
	fmt.Println(obj.foo)

}

type newStruct struct {
	foo int
}
