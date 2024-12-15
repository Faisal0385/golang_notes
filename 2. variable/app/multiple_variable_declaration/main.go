package main

import "fmt"

func main() {
	// Syntax:
	// var var1, var2, ... type
	// var var1, var2 = value1, value2

	// Example:
	var x, y, z int                 // Declares three int variables with a zero value (0)
	var name, age = "Bob", 30       // Type inferred for each
	fmt.Println(x, y, z, name, age) // Output: 0 0 0 Bob 30

	// Using :=
	a, b := 10, 20
	fmt.Println(a, b) // Output: 10 20
}
