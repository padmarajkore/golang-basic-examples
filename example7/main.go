package main

import (
	"fmt"
	"time"
)

// Example 7: Structs
// Demonstrates working with structs (custom data types) in Go

// Basic struct definition
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Struct with embedded struct
type Employee struct {
	Person     // Embedded struct
	EmployeeID string
	Title      string
	Department string
	HireDate   time.Time
}

// Struct with tags for serialization
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price"`
	SKU         string  `json:"sku,omitempty"`
}

// Method for Person struct
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// Method with pointer receiver
func (p *Person) Birthday() {
	p.Age++
}

// Method for Employee struct
func (e Employee) EmployeeInfo() string {
	return fmt.Sprintf("%s, %s - %s", e.FullName(), e.Title, e.Department)
}

func main() {
	// Creating and using structs
	fmt.Println("Structs in Go:")

	// Basic struct initialization
	p1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	fmt.Println("Person:", p1)

	// Short initialization (not recommended - order dependent)
	p2 := Person{"Jane", "Smith", 28}
	fmt.Println("Person 2:", p2)

	// Partial initialization
	p3 := Person{FirstName: "Alice"}
	fmt.Println("Person 3:", p3)

	// Zero value initialization
	var p4 Person
	p4.FirstName = "Bob"
	p4.LastName = "Brown"
	p4.Age = 35
	fmt.Println("Person 4:", p4)

	// Using a method
	fmt.Println("Full name:", p1.FullName())

	// Using a pointer receiver method
	fmt.Printf("Before birthday: %s is %d years old\n", p1.FullName(), p1.Age)
	p1.Birthday()
	fmt.Printf("After birthday: %s is %d years old\n", p1.FullName(), p1.Age)

	// Embedded struct
	employee := Employee{
		Person: Person{
			FirstName: "Sarah",
			LastName:  "Johnson",
			Age:       32,
		},
		EmployeeID: "E12345",
		Title:      "Software Engineer",
		Department: "Engineering",
		HireDate:   time.Date(2018, time.March, 15, 0, 0, 0, 0, time.UTC),
	}

	// Access fields from embedded struct
	fmt.Println("\nEmbedded structs:")
	fmt.Println("Employee Name:", employee.FirstName, employee.LastName)
	fmt.Println("Employee Age:", employee.Age)
	fmt.Println("Employee ID:", employee.EmployeeID)
	fmt.Println("Title:", employee.Title)
	fmt.Println("Department:", employee.Department)
	fmt.Println("Hire Date:", employee.HireDate.Format("2006-01-02"))

	// Access methods from embedded struct
	fmt.Println("Employee Full Name:", employee.FullName())

	// Use struct method
	fmt.Println("Employee Info:", employee.EmployeeInfo())

	// Anonymous struct
	fmt.Println("\nAnonymous struct:")
	point := struct {
		X int
		Y int
	}{10, 20}
	fmt.Println("Point:", point)

	// Comparing structs
	fmt.Println("\nComparing structs:")
	p5 := Person{FirstName: "John", LastName: "Doe", Age: 30}
	fmt.Println("p1 == p5:", p1 == p5)

	// Struct pointers
	fmt.Println("\nStruct pointers:")
	pPtr := &Person{FirstName: "Mark", LastName: "Wilson", Age: 40}
	fmt.Println("Person pointer:", pPtr)
	fmt.Println("Dereferenced:", *pPtr)

	// No need to dereference when accessing fields
	pPtr.Age = 41 // Same as (*pPtr).Age = 41
	fmt.Println("After modifying age:", pPtr)
}