package main

import (
	"fmt"
)

func main() {

	// bool type

	var flag bool = true
	fmt.Printf("%v, %T\n", flag, flag)

	m := 1 == 1
	n := 1 == 2
	fmt.Printf("%v, %v\n", m, n)

	var flag2 bool
	fmt.Printf("%v, %T\n", flag2, flag2)

	// int and uint types

	var a int8 = 0
	var b int16 = 0
	var c int32 = 0
	var d int64 = 0
	var e uint8 = 0
	var f uint16 = 0
	var g uint32 = 0
	var h uint64 = 0
	fmt.Printf("%T, %T, %T, %T, %T, %T, %T, %T\n", a, b, c, d, e, f, g, h)

	// An int divided by an int gives us another int

	var1 := 5
	var2 := 2
	res := var1/var2
	fmt.Println(res)

	// Bit operations on integers

	p := 10 // 1010
	q := 3 // 0011
	fmt.Println(p & q) // 0010 (and)
	fmt.Println(p | q) // 1011 (or)
	fmt.Println(p ^ q) // 1001 (xor)
	fmt.Println(p &^ q) // 0100 (and not)

	// Bit-shifting integers

	bit := 8 // 2^3
	fmt.Println(bit << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(bit >> 3) // 2^3 / 2^3 = 2^0

	// float type

	float := 3.14
	float = 1.1e10
	float = 1.23E30
	fmt.Println(float)

	// string type

	str := "string"
	bytes := []byte(str)
	fmt.Printf("%v, %T\n", bytes, bytes)

	// rune type

	var r rune = 'آ'
	fmt.Printf("%v, %T\n", r, r)
}