package main // Entering point of application

import "fmt"

func main() {

	// Calling a function in package

	saySomething("Hello, world!")

	saySomethings("Hello again, world!", 10)

	addUp(10, 13)

	// Swapping two variables

	a := 10
	b := 5

	fmt.Println(a, b)

	swap(&a, &b)

	fmt.Println(a, b)

	// Using variadic parameter function

	printSum(10)
	printSum(1, 2, 3, 4, 5, 6)
	printSum(100, 200, 300, 400, 500, 600, 700, 800, 900)

	fmt.Println(returnSum(1, 2, 3))

	fmt.Println(returnSumNew(1, 2, 3, 4))

	fmt.Println(*returnSumPointer(1, 2, 3, 4, 5) + 1)

	// Using a function that returns multiple parameters

	c, err := divide(2, 0)

	if err == nil {
		fmt.Println(c)
	}

	// Anonymous function

	func() {
		fmt.Println("anonymous")
	}()

	/*
	For avoiding odd behavior using anonymous functions in syncs, instead of this code:

	for i := 0; i < 10; i++ {
		func() {
			fmt.Println(i)
		}()
	}

	We'll use the code below:
	 */

	for i := 0; i < 10; i++ {
		func(num int) {
			fmt.Println(num)
		}(i)
	}

	// Define function as a variable

	f := func(i int) {
		fmt.Println("defined function, input number:", i)
	}

	f(1)
	f(2)
	f(3)

	// Using a method

	sina := user{
		name: "sina",
	}

	sina.greetUser()

	fmt.Println(sina.name)
	sina.setUserName("sina2")
	fmt.Println(sina.name)

}

// Void function

func saySomething (message string)  {
	fmt.Println(message)
}

// Multi-parameter function

func saySomethings (message string, times int)  {
	for i := 0; i < times; i++ {
		fmt.Println(message)
	}
}

func addUp (a, b int)  {
	fmt.Println(a + b)
}

// Function with pointers

func swap (px, py *int) {
	temp := *px
	*px = *py
	*py = temp
}

/*
 Variadic parameter function (with a number of variables that is not defined, like "*args" and "**kwargs" in Python)
 Variadic parameters must be the last parameters in function; For example, the function below, can be re-written like
 "func printSum (str string, numbers ...int)" but not "func printSum (numbers ...int, str string)"/
 */

func printSum (numbers ...int) {
	res := 0
	for _, number := range numbers {
		res += number
	}
	fmt.Printf("Sum of the numbers: %v\n", res)
}

// Returning values

func returnSum (numbers ...int) int {
	res := 0
	for _, number := range numbers {
		res += number
	}
	return res
}

// Named return values

func returnSumNew (numbers ...int) (res int) {
	for _, number := range numbers {
		res += number
	}
	return
}

// Returning pointer values

func returnSumPointer (numbers ...int) *int {
	res := 0
	for _, number := range numbers {
		res += number
	}
	return &res
}

// Returning a parameter and an error

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("division by zero")
	}
	return a/b, nil
}

// Methods in Go

type user struct {
	name string
}

func (u user) greetUser() {
	fmt.Printf("Hello, %v\n", u.name)
}

func (u *user) setUserName(str string) {
	u.name = str
}