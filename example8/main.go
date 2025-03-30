package main

import (
	"fmt"
	"math"
)

// Example 8: Interfaces
// Demonstrates working with interfaces in Go

// Shape interface defines a method for calculating area
type Shape interface {
	Area() float64
}

// Additional interface that some shapes might implement
type Perimeter interface {
	Perimeter() float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

// Area method for Circle - implements Shape interface
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter method for Circle - implements Perimeter interface
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area method for Rectangle - implements Shape interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter method for Rectangle - implements Perimeter interface
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area method for Triangle - implements Shape interface
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Function that takes a Shape interface
func printArea(s Shape) {
	fmt.Printf("Area of %T: %.2f\n", s, s.Area())
}

// Function that takes both interfaces
func printShapeInfo(s interface{}) {
	fmt.Printf("Shape type: %T\n", s)

	// Type assertion
	if shape, ok := s.(Shape); ok {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}

	// Type assertion for Perimeter
	if perim, ok := s.(Perimeter); ok {
		fmt.Printf("Perimeter: %.2f\n", perim.Perimeter())
	} else {
		fmt.Println("This shape doesn't implement Perimeter")
	}

	fmt.Println() // Add a blank line
}

// A completely different type that also implements Shape
type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

// A method that demonstrates type switches
func classifyShape(item interface{}) {
	switch v := item.(type) {
	case Circle:
		fmt.Printf("Circle with radius: %.2f\n", v.Radius)
	case Rectangle:
		fmt.Printf("Rectangle with width: %.2f and height: %.2f\n", v.Width, v.Height)
	case Triangle:
		fmt.Printf("Triangle with base: %.2f and height: %.2f\n", v.Base, v.Height)
	case Square:
		fmt.Printf("Square with side: %.2f\n", v.Side)
	case Shape:
		fmt.Println("Some other shape with area:", v.Area())
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

// Empty interface example function - accepts any value
func printAny(v interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", v, v)
}

func main() {
	// Create shapes
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 3, Height: 4}
	triangle := Triangle{Base: 6, Height: 8}
	square := Square{Side: 4}

	// Using the Shape interface
	fmt.Println("Using the Shape interface:")
	printArea(circle)
	printArea(rectangle)
	printArea(triangle)
	printArea(square)

	// Using type assertions and multiple interfaces
	fmt.Println("\nShape information using type assertions:")
	printShapeInfo(circle)
	printShapeInfo(rectangle)
	printShapeInfo(triangle)
	printShapeInfo(square)

	// Using type switches
	fmt.Println("Using type switches:")
	shapes := []interface{}{
		circle,
		rectangle,
		triangle,
		square,
		"Not a shape", // Will go to default case
		42,            // Will go to default case
	}

	for _, shape := range shapes {
		classifyShape(shape)
	}

	// Using the empty interface
	fmt.Println("\nEmpty interface examples:")
	printAny(42)
	printAny("Hello, interfaces!")
	printAny(true)
	printAny(circle)

	// Creating a slice of shapes (using the interface as a type)
	fmt.Println("\nSlice of Shapes:")
	shapes2 := []Shape{circle, rectangle, triangle, square}

	// Calculate the total area
	totalArea := 0.0
	for _, shape := range shapes2 {
		totalArea += shape.Area()
	}
	fmt.Printf("Total area of all shapes: %.2f\n", totalArea)
}