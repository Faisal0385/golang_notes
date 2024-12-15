package main

import "fmt"

func main() {
	// Syntax:
	// var (
	//     var1 type = value
	//     var2 type = value
	// )

	// const (
	//     const1 = value
	//     const2 = value
	// )

	// Example:

	var (
		name = "Charlie"
		age  = 35
	)

	const (
		pi       = 3.14
		language = "Go"
	)

	fmt.Println(name, age, pi, language) // Output: Charlie 35 3.14 Go
}
