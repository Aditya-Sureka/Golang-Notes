// The range form of the for loop iterates over a slice, map, array, string, or channel.
// When ranging over an array, slice, string, or map, the index and value are returned for each iteration. When ranging over a channel, only the value is returned until the channel is closed.
// The range expression is evaluated once before the loop starts, so if you pass a reference to the iteration variable (e.g., &num), it will point to the same memory address in each iteration, which can lead to unexpected behavior. To avoid this, you can use a local variable within the loop to capture the current value of the iteration variable.
// When ranging over a string, the index and the Unicode code point of each character are returned. The index is the byte position of the character in the string, and the value is the Unicode code point (rune) of the character. To convert the rune to a string, you can use the string() function.
// When ranging over a map, the order of iteration is not specified and can vary from one run to another. If you need a specific order, you can extract the keys into a slice, sort it, and then iterate over the sorted keys to access the corresponding values in the map.
// When ranging over a channel, the loop will continue until the channel is closed. If you need to break out of the loop based on a condition, you can use a select statement with a case for receiving from the channel and another case for a timeout or other condition.
// The range form of the for loop is a powerful and convenient way to iterate over various data structures in Go, providing both the index and value for each element, and it can be used with arrays, slices, maps, strings, and channels.
// For more details, see the official Go documentation on the range clause: https://golang.org/ref/spec#RangeClause

package main

import "fmt"

func main() {

	// Range with arrays and slices
	nums := []int{6, 7, 8}
	for i, num := range nums {
		fmt.Println(i, num)
	}

	// Range with maps
	m := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range m {
		fmt.Println(key, value)
	}

	// Range with strings
	for i, c := range "golang" {
		fmt.Println(i, string(c)) // c is the Unicode code point of the character
	}
}