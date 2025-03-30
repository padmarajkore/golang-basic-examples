package main

import (
	"fmt"
	"sort"
)

// Example 6: Maps
// Demonstrates working with maps (key-value pairs) in Go
func main() {
	// Creating a map
	fmt.Println("Maps in Go:")

	// Declare an empty map
	var emptyMap map[string]int
	fmt.Println("Empty map (nil):", emptyMap)

	// Initialize using make
	ages := make(map[string]int)

	// Add key-value pairs
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 22
	fmt.Println("Ages map:", ages)

	// Map literal
	colors := map[string]string{
		"red":    "#FF0000",
		"green":  "#00FF00",
		"blue":   "#0000FF",
		"yellow": "#FFFF00",
	}
	fmt.Println("Colors map:", colors)

	// Accessing values
	aliceAge := ages["Alice"]
	fmt.Println("Alice's age:", aliceAge)

	// Accessing non-existent key
	// This will return the zero value for the value type (0 for int)
	unknownAge := ages["Unknown"]
	fmt.Println("Unknown's age:", unknownAge)

	// Check if a key exists
	age, exists := ages["Bob"]
	if exists {
		fmt.Println("Bob's age exists:", age)
	}

	age, exists = ages["David"]
	if !exists {
		fmt.Println("David's age doesn't exist")
	}

	// Delete a key-value pair
	delete(ages, "Charlie")
	fmt.Println("After deleting Charlie:", ages)

	// Length of a map
	fmt.Println("Number of entries in colors:", len(colors))

	// Iterating over a map
	fmt.Println("\nIterating over the colors map:")
	for key, value := range colors {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Note: Maps are unordered, but we can sort the keys
	fmt.Println("\nIterating in sorted key order:")
	keys := make([]string, 0, len(colors))
	for k := range colors {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("Key: %s, Value: %s\n", k, colors[k])
	}

	// Maps as reference types
	fmt.Println("\nMaps as reference types:")
	original := map[string]int{"one": 1, "two": 2}
	fmt.Println("Original:", original)

	// Creating a reference, not a copy
	reference := original
	reference["three"] = 3

	fmt.Println("Original after modifying reference:", original)
	fmt.Println("Reference:", reference)

	// Making a true copy
	copy := make(map[string]int)
	for k, v := range original {
		copy[k] = v
	}
	copy["four"] = 4

	fmt.Println("Original after modifying copy:", original)
	fmt.Println("Copy:", copy)

	// Maps of maps (nested maps)
	fmt.Println("\nNested maps:")
	cityPopulations := map[string]map[string]int{
		"USA": {
			"New York":    8419000,
			"Los Angeles": 3979576,
			"Chicago":     2699000,
		},
		"Japan": {
			"Tokyo": 9273000,
			"Osaka": 2691000,
			"Kyoto": 1459000,
		},
	}

	// Access nested map
	fmt.Println("Population of New York:", cityPopulations["USA"]["New York"])

	// Add a city to an existing country
	cityPopulations["USA"]["Boston"] = 675647

	// Add a new country with cities
	cityPopulations["Canada"] = map[string]int{
		"Toronto":   2930000,
		"Montreal":  1780000,
		"Vancouver": 675218,
	}

	// Print all populations
	fmt.Println("\nAll city populations:")
	for country, cities := range cityPopulations {
		fmt.Println(country + ":")
		for city, population := range cities {
			fmt.Printf("  %s: %d\n", city, population)
		}
	}
}