package main

import (
	"fmt"
)

// Example 4: Functions
// Demonstrates various function types and usages in Go

// Basic function with parameters and return value
func add(a, b int) int {
	return a + b
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Function with named return values
func getMinMax(numbers []int) (min, max int) {
	min = numbers[0]
	max = numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	// Return statement will automatically return named return values
	return
}

// Variadic function (variable number of arguments)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function with a function parameter (higher-order function)
func applyAndPrint(numbers []int, f func(int) int) {
	for _, num := range numbers {
		fmt.Printf("f(%d) = %d\n", num, f(num))
	}
}

// Closure (function that references variables outside its body)
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Defer example function
func deferExample() {
	fmt.Println("Starting deferExample()")
	defer fmt.Println("This will print at the end")
	defer fmt.Println("This will print second-to-last")
	fmt.Println("Middle of deferExample()")
}

func main() {
	// Basic function
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)

	// Multiple return values
	result2, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result2)
	}

	result3, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 0 =", result3)
	}

	// Named return values
	numbers := []int{7, 2, 9, 3, 6}
	min, max := getMinMax(numbers)
	fmt.Printf("In %v, min is %d and max is %d\n", numbers, min, max)

	// Variadic function
	fmt.Println("Sum of 1, 2, 3, 4, 5:", sum(1, 2, 3, 4, 5))

	// Higher-order function
	fmt.Println("Applying functions to numbers:")
	square := func(x int) int { return x * x }
	applyAndPrint([]int{1, 2, 3, 4}, square)

	// Closure
	counter := makeCounter()
	fmt.Println("Counter:", counter())
	fmt.Println("Counter:", counter())
	fmt.Println("Counter:", counter())

	// Defer
	deferExample()
}