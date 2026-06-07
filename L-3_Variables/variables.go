/* Code for Writing Variables in Go
   In Go, you can declare variables in several ways
   In Go, whenever you declare a variable, you must specify its type
   In Go, whenever you declare a variable, you must also call it somewhere otherwise it will result in a compile-time error
*/

package main

import "fmt"

func main() {
	var name string = "golang" // Declaring a variable of type string and assigning a value
	fmt.Println(name)  // Output: golang

	var new = "hello"; // Declaring a variable using type inference, Go will automatically determine the type based on the assigned value
	fmt.Println(new) // Output: hello

	var isAvailable bool; // Declaring a boolean variable without initializing it, it will have the zero value of false
	fmt.Println(isAvailable) // Output: false

	var isAdult = true; // Declaring a boolean variable using type inference
	fmt.Println(isAdult) // Output: true

	var age int; // Declaring an integer variable without initializing it, it will have the zero value of 0
	fmt.Println(age) // Output: 0

	/* Shorthand Syntax */
	// In Go, you can also declare and initialize a variable using shorthand syntax, which is a more concise way to declare variables. The syntax is as follows:
	// variableName := value
	// This syntax is only allowed within functions and cannot be used at the package level.

	intro := "golang" // Declaring and initializing a variable using shorthand syntax
	fmt.Println(intro) // Output: golang

	output := true // Declaring and initializing a boolean variable using shorthand syntax
	fmt.Println(output) // Output: true
}