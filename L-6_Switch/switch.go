/* Code for Switch Statements */
// Switch statements are a more concise way to write multiple if-else statements. They allow you to compare a variable against multiple values and execute different code based on the match.
// In Go, switch statements can be used with various types, including integers, strings, and even custom types. The switch statement evaluates the expression and compares it against the cases. If a match is found, the corresponding code block is executed. If no match is found, the default case is executed (if provided).
// Here's an example of a switch statement in Go:
// In this example, we have a variable `i` with the value 1. The switch statement checks the value of `i` against the cases. Since `i` is equal to 1, it matches the first case and prints "one". If `i` were 2, it would print "two". If `i` were any other value, it would print "other" due to the default case.
// You can also use switch statements without an expression, which allows you to evaluate multiple conditions in a more flexible way. For example:
/*
switch {
case i < 0:
	fmt.Println("negative")
case i == 0:
	fmt.Println("zero")
case i > 0:
	fmt.Println("positive")
}
*/
// In this example, the switch statement evaluates the conditions directly without an expression. It checks if `i` is negative, zero, or positive and prints the corresponding message. This can be useful when you want to evaluate multiple conditions that are not based on a single variable.

package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")
	}

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("Its an string")
		case bool:
			fmt.Println("Its an boolean")
		default: 
		fmt.Println("other", t)
		}
	}
	whoAmI(55.56)
}