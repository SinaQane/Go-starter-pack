package main

import (
	"fmt"
	"reflect"
)

// Creating a struct

type User struct {
	id int `required min:"18"`
	name string
	phoneNumber string
	email string
}

// Inheritance in go (We can't extend in go, we can just use some kind of composition)

type Animal struct {
	name string
	age int
}

type Bird struct {
	Animal
	canFly bool
}

func main() {

	// Creating an object from a struct type

	sina := User{
		id: 0,
		name: "sina",
		phoneNumber: "0939",
		email: "qanqan@qan.qan",
	}

	fmt.Println(sina)
	fmt.Println(sina.email)

	// Dumb way to create an object from a struct (Don't use it)

	hirbod := User{
		1,
		"hirbod",
		"0912",
		"hirhir@hir.hir",
	}

	fmt.Println(hirbod)

	// Anonymous struct

	maryam := struct {
		name string
	}{name: "maryam"}

	fmt.Println(maryam)

	// Structs are not pointer reference by default

	maryam2 := maryam
	maryam2.name = "maryam2"

	fmt.Println(maryam)
	fmt.Println(maryam2)

	// Create an object from a struct that uses inheritance

	b := Bird{}
	b.name = "Crow"
	b.age = 10
	b.canFly = true

	fmt.Println(b)

	// Tag

	tag := reflect.TypeOf(User{})
	field, _ := tag.FieldByName("id")
	fmt.Println(field.Tag)
}
