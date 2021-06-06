package main

import (
	"bytes"
	"fmt"
	"io"
)

func main()  {

	// Using interfaces

	var w Typer = MyTyper{}
	_, err := w.Type([]byte("Hello, World!"))
	if err != nil {
		return
	}

	myCounter := Counter(0)
	var add Adder = &myCounter
	for i := 0; i < 10; i++ {
		fmt.Println(add.Add())
	}

	var wc WriterCloser = NewBufferedWriterCloser()
	_, err = wc.Write([]byte("Hello, World!"))
	if err != nil {
		return
	}
	err = wc.Close()
	if err != nil {
		return
	}

	// Casting

	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)

	// Casting with error handling

	// The code below will cause to a panic. since WriterCloser doesn't have a Read method

	/*
	r := wc.(io.Reader)
	fmt.Println(r)
	*/

	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("conversion failed")
	}

	// Using empty interfaces

	var obj interface{} = NewBufferedWriterCloser()
	if wc, ok := obj.(WriterCloser); ok {
		_, err := wc.Write([]byte("Hello, World!"))
		if err != nil {
			return
		}
	} else {
		fmt.Println("empty interface conversion failed")
	}

}

// Simple interface example (using struct to implement interface)

type Typer interface {
	Type([]byte) (int, error)
}

type MyTyper struct {}

func (w MyTyper) Type(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// Using custom type to implement interface

type Adder interface {
	Add() int
}

type Counter int

func (cnt *Counter) Add() int {
	*cnt ++
	return int(*cnt)
}

// Implementing more than one interface

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

	// I was too lazy to write the whole code here, so I just left them empty...

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func (bwc *BufferedWriterCloser) Close() error {
	return nil
}

	// Constructor for BufferedWriterCloser

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
