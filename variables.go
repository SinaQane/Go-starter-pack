package main

import (
	"fmt"
	"strconv"
)

// Package level visible variable (first letter lowercase)

var variable int = 80

// Global level visible variable (first letter uppercase)

var Variable int = 80

// var block for more clean code

var (
	string1 string = "hi"
	string2 string = "hello"
	string3 string = "bonjour"
)

// Declaring a variable that we'll declare again in main function scope

var twiceDeclaration int = 100

func main() {

	// Hello World!

	fmt.Println("hello world")

	// Declaring a variable

	var i int
	i = 20

	var j int = 40

	k := 60. // float64

	fmt.Printf("%v + %v + %v Types: %T, %T, %T\n", i, j, k, i, j, k)

	// Declaring a variable in package scope and in function scope another time

	twiceDeclaration := 40  // Works fine even tho it's been declared once before
	fmt.Printf("%v\n", twiceDeclaration)

	// Convert int to float

	var oldK int = 40
	fmt.Printf("%v, %T\n", oldK, oldK)

	newK := float32(oldK)
	fmt.Printf("%v, %T\n", newK, newK)

	// Convert int to string

	var str1 int = 10
	fmt.Printf("%v, %T\n", str1, str1)

	var str2 string = string(str1) // Doesn't work
	fmt.Printf("%v, %T\n", str2, str2)

	var str3 string = strconv.Itoa(str1) // Works
	fmt.Printf("%v, %T\n", str3, str3)
}