package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Example 10: Error Handling
// Demonstrates error handling techniques in Go

// Standard error handling with return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Custom error type
type ValidationError struct {
	Field   string
	Message string
}

// Implementation of the Error interface for custom error type
func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on field %s: %s", e.Field, e.Message)
}

// Function that returns our custom error type
func validatePositive(value int, name string) error {
	if value <= 0 {
		return ValidationError{
			Field:   name,
			Message: "must be positive",
		}
	}
	return nil
}

// Function to demonstrate wrapping errors
func parseAndMultiply(str string, factor int) (int, error) {
	// Parse string to int
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("error parsing '%s': %w", str, err)
	}

	// Validate the number
	if err := validatePositive(num, "parsed number"); err != nil {
		return 0, fmt.Errorf("validation failed: %w", err)
	}

	// Validate the factor
	if err := validatePositive(factor, "factor"); err != nil {
		return 0, fmt.Errorf("factor validation failed: %w", err)
	}

	return num * factor, nil
}

// Deferred function that handles cleanup
func readFile(path string) (string, error) {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	// Important: defer the close operation to ensure it happens even if
	// an error occurs later
	defer file.Close()

	// Read the file content
	var sb strings.Builder
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return "", fmt.Errorf("failed to read file: %w", err)
		}

		if n == 0 {
			break
		}

		sb.Write(buf[:n])
	}

	return sb.String(), nil
}

// Panic and recover demonstration
func safeDivision(a, b int) (result int) {
	// Set up deferred recovery function
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			result = 0 // Set a default value
		}
	}()

	// This will panic if b is 0
	if b == 0 {
		panic("division by zero")
	}

	return a / b
}

// Function that creates a temporary file and ensures cleanup
func workWithTempFile() error {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "example")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}

	// Ensure we clean up the file even if an error occurs
	defer func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
		fmt.Println("Temporary file cleaned up:", tempFile.Name())
	}()

	// Write something to the file
	_, err = tempFile.Write([]byte("This is a temporary file."))
	if err != nil {
		return fmt.Errorf("failed to write to temp file: %w", err)
	}

	fmt.Println("Successfully wrote to temporary file:", tempFile.Name())
	return nil
}

func main() {
	fmt.Println("Error Handling in Go:")

	// Basic error handling
	fmt.Println("\n=== Basic Error Handling ===")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of 10/2:", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of 10/0:", result) // This won't execute
	}

	// Custom error types
	fmt.Println("\n=== Custom Error Types ===")
	err = validatePositive(5, "count")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Validation passed for 5")
	}

	err = validatePositive(-3, "amount")
	if err != nil {
		fmt.Println("Error:", err)

		// Type assertion
		if valErr, ok := err.(ValidationError); ok {
			fmt.Printf("  Field: %s, Message: %s\n", valErr.Field, valErr.Message)
		}
	}

	// Error wrapping
	fmt.Println("\n=== Error Wrapping ===")
	product, err := parseAndMultiply("10", 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", product)
	}

	product, err = parseAndMultiply("abc", 5)
	if err != nil {
		fmt.Println("Error:", err)
	}

	product, err = parseAndMultiply("-5", 3)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Defer for cleanup
	fmt.Println("\n=== Defer for Cleanup ===")
	// This will likely fail, but cleanup will still happen
	content, err := readFile("nonexistent_file.txt")
	if err != nil {
		fmt.Println("Read file error:", err)
	} else {
		fmt.Println("File content length:", len(content))
	}

	// Try to work with a temporary file
	err = workWithTempFile()
	if err != nil {
		fmt.Println("Temp file error:", err)
	}

	// Panic and recover
	fmt.Println("\n=== Panic and Recover ===")
	fmt.Println("Safe division 10/2:", safeDivision(10, 2))
	fmt.Println("Safe division 10/0:", safeDivision(10, 0))

	fmt.Println("\nError handling examples completed")
}