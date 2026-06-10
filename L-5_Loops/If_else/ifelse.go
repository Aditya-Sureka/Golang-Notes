/* Code for If-Else Statements */
// If-else statements are used to execute a block of code based on a condition.

package main

import "fmt"

func main() {
	// If-else loops -> 
	age := 16

	if age >= 18 {
		fmt.Println("Person is Adult")
	} else {
		fmt.Println("Person is not an adult")
	}

	// Else-if loopps -> 
	price := 14

	if price >= 18 {
		fmt.Println("Price is greater than or equal to 18")
	} else if price >= 12 {
		fmt.Println("Price is greater than 12 but less than 18")
	} else {
		fmt.Println("Price is less than 18")
	}

	// Logical operators -> 
	var role = "admin"
	var hasAccess = false

	if role == "admin" &&  hasAccess {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	} 

	// We can declare a variable in the if statement and use it in the else-if and else block as well.
	if age := 15; age >= 18 {
		fmt.Println("person is adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager", age)
	}

	/* TERNARY OPERATOR -> Go does not have a ternary operator like other languages, but we can achieve similar functionality using if-else statements. */
	var number = 10
	var result string
	if number%2 == 0 {
		result = "Even"
	} else {
		result = "Odd"
	}
	fmt.Println("The number is:", result)
}