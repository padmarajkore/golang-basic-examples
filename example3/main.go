package main

import "fmt"

// Example 3: Control Flow
// Demonstrates if/else, loops, and switch statements
func main() {
	// If-else statement
	age := 18
	if age >= 18 {
		fmt.Println("You can vote!")
	} else {
		fmt.Println("You cannot vote yet.")
	}

	// If with short statement
	if score := 85; score >= 70 {
		fmt.Println("You passed with a score of", score)
	} else {
		fmt.Println("You failed with a score of", score)
	}

	// For loop - standard
	fmt.Println("Standard for loop:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// For loop - as a while loop
	fmt.Println("While-style loop:")
	count := 0
	for count < 3 {
		fmt.Println(count)
		count++
	}

	// For loop - infinite with break
	fmt.Println("Loop with break:")
	sum := 0
	for {
		sum++
		if sum > 5 {
			break
		}
		fmt.Println("Sum:", sum)
	}

	// For loop - with continue
	fmt.Println("Loop with continue (even numbers only):")
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	// Switch statement
	day := "Wednesday"
	fmt.Println("Switch example:")
	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Wednesday":
		fmt.Println("Middle of the work week")
	case "Friday":
		fmt.Println("End of the work week")
	default:
		fmt.Println("Some other day")
	}

	// Switch with no expression (like if-else-if)
	fmt.Println("Switch with no expression:")
	hour := 15
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}