/* Code for the For loops section of the Go course. */
/* In Go, there is only one looping construct: the for loop. */
/* The for loop has three components: the initialization, the condition, and the post statement. */
/* The initialization is executed before the loop starts. It is typically used to initialize a counter variable. */
/* The condition is evaluated before each iteration of the loop. If the condition is true, the loop body is executed. If the condition is false, the loop terminates. */
/* The post statement is executed after each iteration of the loop. It is typically used to update the counter variable. */
/* The syntax of the for loop is as follows: */

/*
for initialization; condition; post {
	// loop body
}
*/

package main

import "fmt"

func main() {
	// while loop using for
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// for loop with all three components
	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	// infinite loop
	for {
	 	fmt.Println("This is an infinite loop.")
	}

	/* Range Loops -> 
	// The range keyword is used to iterate over elements in a variety of data structures,
	// such as arrays, slices, maps, and strings. It provides a convenient way to access both 
	// the index and the value of each element in the data structure.
	*/

	for k := range 3 {
		fmt.Println(k)
	}
}