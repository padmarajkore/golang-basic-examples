package main

import "fmt"

// Example 5: Arrays and Slices
// Demonstrates working with arrays and slices in Go
func main() {
	// Arrays
	fmt.Println("Arrays:")

	// Array declaration with fixed size
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Println("Array:", numbers)

	// Array initialization with values
	fruits := [3]string{"apple", "banana", "orange"}
	fmt.Println("Fruits:", fruits)

	// Array with inferred size
	colors := [...]string{"red", "green", "blue", "yellow"}
	fmt.Println("Colors:", colors)
	fmt.Println("Number of colors:", len(colors))

	// 2D array
	matrix := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Matrix:", matrix)

	// Iterating through an array
	fmt.Println("Iterating through fruits:")
	for i, fruit := range fruits {
		fmt.Printf("fruits[%d] = %s\n", i, fruit)
	}

	// Slices
	fmt.Println("\nSlices:")

	// Creating a slice from an array
	numbersSlice := numbers[1:4] // items 1 through 3
	fmt.Println("Slice from array:", numbersSlice)

	// Changing a slice affects the underlying array
	numbersSlice[0] = 25
	fmt.Println("After modification - Slice:", numbersSlice)
	fmt.Println("After modification - Original array:", numbers)

	// Creating a slice with literals
	letters := []string{"a", "b", "c", "d"}
	fmt.Println("Letters slice:", letters)

	// Creating a slice with make
	scores := make([]int, 5) // len=5, cap=5
	fmt.Println("Scores (make):", scores)

	scores2 := make([]int, 3, 5) // len=3, cap=5
	fmt.Printf("Scores2: %v, Length: %d, Capacity: %d\n", scores2, len(scores2), cap(scores2))

	// Appending elements to a slice
	vowels := []string{"a", "e"}
	fmt.Println("Original vowels:", vowels)

	vowels = append(vowels, "i")
	vowels = append(vowels, "o", "u")
	fmt.Println("After append:", vowels)

	// Slice capacity and growth
	s := make([]int, 0)
	fmt.Printf("Empty slice - Length: %d, Capacity: %d\n", len(s), cap(s))

	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("After adding %d - Length: %d, Capacity: %d\n", i, len(s), cap(s))
	}

	// Slicing operations
	allLetters := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println("All letters:", allLetters)
	fmt.Println("First three:", allLetters[:3])
	fmt.Println("Last three:", allLetters[len(allLetters)-3:])
	fmt.Println("Middle:", allLetters[2:5])

	// Copy slices
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copied := copy(dst, src)
	fmt.Printf("Src: %v, Dst: %v, Copied: %d\n", src, dst, copied)
}