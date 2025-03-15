package main

import (
	"fmt"
	"math/rand"
)

func Session1() {
	fmt.Println("-*- Day 1 -*-")
	fmt.Println("My favorite number is:", add(rand.Intn(10), rand.Intn(20)))

	// variable
	var sum, diff int
	var a, b int = 10, 12

	sum, diff = sum_and_diff(a, b)
	fmt.Printf("The value of func sum_and_diff: %d and %d \n", sum, diff)

	m := multiply(a, b)
	fmt.Printf("The value of func multiply: %v \n", m)
	fmt.Println("--- *** ---")
}

// function:
func add(a, b int) int {
	return a + b
}

// `Naked` return:
func sum_and_diff(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

// Short variable
func multiply(a, b int) int {
	m := a * b
	return m
}
