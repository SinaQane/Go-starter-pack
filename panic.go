package main

import (
	"fmt"
	"net/http"
)

func main() {

	// panic is like "throw" syntax in other languages, here are some examples:

	// Example 1

	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)

	// Example 2

	fmt.Println("hi")
	panic("panicking")
	fmt.Println("bye")

	// Example 3

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		panic(err.Error())
	}

	// Example 4

	fmt.Println("hi")
	defer 	fmt.Println("first defer")
	panic("panicking")
	fmt.Println("bye")
	defer 	fmt.Println("second defer")
}
