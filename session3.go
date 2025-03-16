package main

import (
	"fmt"
)

func Session3() {
	fmt.Println()
	fmt.Println("-*- Day 3 -*-")
	pointerExample()
	structExample()
}

// Pointers
func pointerExample() {
	fmt.Println("Go has pointers, but no pointer arithmetic")
	i := 5
	p := &i
	fmt.Printf("Go using `&` to get address of parameter: p := &i => p=%v \n", p)
	fmt.Printf("Using `*` show the underlying value of => *p=%v \n", *p)
}

// Struct
// Like C++, Go has Struct.
type Person struct {
	name string
	age  int
}

func structExample() {
	p1 := Person{name: "Dotoki", age: 24}
	fmt.Println("Example about Struct:")
	fmt.Printf("Name: %v | Age: %v\n", p1.name, p1.age)
}
