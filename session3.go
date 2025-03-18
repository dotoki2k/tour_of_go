package main

import (
	"fmt"
)

func Session3() {
	fmt.Println()
	fmt.Println("-*- Day 3 -*-")
	pointerExample()
	structExample()
	arrayExample()
	slicesExample()
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

// Arrays
func arrayExample() {
	var a [4]string
	a[0] = "Hello"
	a[1] = "Array"
	a[2] = a[0]
	a[3] = "Dotoki"
	fmt.Println(a)
	fmt.Println(a[2:])
}

// Slices are like references to arrays
func slicesExample() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("The original array:", names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println("Slice a from index 0 to 2:", a)
	fmt.Println("Slice b from index 1 to 3:", b)
	fmt.Println("Change the first element on slice `b` to NoName, that will affect to original array and slice `a`.")
	b[0] = "NoName"
	fmt.Println("Slice a and b: ", a, b)
	fmt.Println(names)
}
