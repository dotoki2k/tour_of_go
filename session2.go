package main

import (
	"fmt"
)

func Session2() {
	fmt.Println("-*- Day 2 -*-")
	fmt.Println(looping())
	fmt.Println(whileLooping())
	fmt.Println(isOdd(2))
	fmt.Println(sqrt(200))
	fSwitch(0)
	exampleDefer()
	fmt.Println("--- *** ---")
}

// Go has Only one looping construct.
func looping() (result string) {
	sum := 0
	for i := 1; i < 11; i++ {
		sum += i
	}
	result = fmt.Sprintf("The sum from 1 to 10 is %v by For", sum)
	return
}

// while in Go :))

func whileLooping() (result string) {
	sum := 0
	i := 0
	for i < 10 {
		i++
		sum += i
	}
	result = fmt.Sprintf("The sum from 1 to 10 is %v by `While` :v", sum)
	return
}

// if
func isOdd(a int) string {
	odd := "not"
	if a%2 != 0 {
		odd = ""
	}
	return fmt.Sprintf("%v is %v odd", a, odd)
}

// exercise: find the number z for which z² is most nearly x

func sqrt(x float64) string {

	var before, z float64
	z = 1
	before = 0
	count := 1
	for ; count < 100; count++ {
		before = z
		z -= (z*z - x) / (2 * z)
		if before == z || z*z == x {
			break
		}
	}
	result := fmt.Sprintf("%v is most nearly after loop %v times.", z, count)
	return result
}

// switch

func fSwitch(a int) {
	fmt.Println("Switch case")
	switch {
	case a%2 == 0:
		fmt.Printf("%v is even\n", a)
	case a%2 == 1:
		fmt.Printf("%v is odd\n", a)
	default:
		fmt.Printf("default case!")
	}
}

// defer
func exampleDefer() {
	fmt.Print("Defer: ")
	defer fmt.Println()
	for i := 0; i < 11; i++ {
		defer fmt.Printf("%v ", i)
	}
	defer fmt.Println("A defer statement defers the execution of a function until the surrounding function returns.\nDeferred function calls are pushed onto a stack")

}
