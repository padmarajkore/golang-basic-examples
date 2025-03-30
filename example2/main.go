package main

import "fmt"

// Example 2: Variables and Types
// Demonstrates different variable types and declarations in Go
func main() {
	// Variable declaration with explicit type
	var age int = 30
	var name string = "John"
	var isEmployed bool = true

	// Short variable declaration
	salary := 50000.50 // float64

	// Constants
	const pi = 3.14159

	// Multiple variable declaration
	var a, b, c int = 1, 2, 3

	// Zero values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	// Print all variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Employed:", isEmployed)
	fmt.Println("Salary:", salary)
	fmt.Println("Pi:", pi)
	fmt.Println("Multiple variables:", a, b, c)
	fmt.Println("Default values:", defaultInt, defaultFloat, defaultString, defaultBool)
}