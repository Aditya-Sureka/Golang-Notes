// Slices are a more powerful and flexible data structure than arrays in Go.
// A slice is a dynamically-sized, flexible view into the elements of an array. It is a reference type that provides a more convenient and efficient way to work with collections of data compared to arrays. Slices are built on top of arrays and provide additional functionality, such as dynamic resizing and the ability to share underlying data between multiple slices.
// In Go, a slice is defined using the []type syntax, where type is the type of the elements in the slice. For example, you can declare a slice of integers like this: var nums []int. This creates a slice that can hold integers but does not allocate any memory for the underlying array yet. You can also initialize a slice with values at the time of declaration using a composite literal, like this: nums := []int{1, 2, 3}. This creates a slice with an underlying array that contains the specified values.
// Slices have a length and a capacity. The length of a slice is the number of elements it currently holds, while the capacity is the total number of elements that can be stored in the underlying array before it needs to be resized. When you append elements to a slice using the built-in append function, if the length exceeds the capacity, Go automatically allocates a new underlying array with a larger capacity and copies the existing elements to it. This allows slices to grow dynamically as needed without requiring manual resizing.
// Slices also support slicing operations, which allow you to create new slices that reference a portion of an existing slice. For example, if you have a slice nums := []int{1, 2, 3, 4, 5}, you can create a new slice that references the first three elements like this: subSlice := nums[:3]. This creates a new slice that shares the same underlying array as nums but has its own length and capacity. Modifying the elements
// in subSlice will also modify the corresponding elements in nums since they share the same underlying array. This is an important aspect of slices to keep in mind when working with them, as it can lead to unintended side effects if you are not careful with how you manipulate slices and their underlying arrays.
// In summary, slices in Go are a powerful and flexible data structure that provides dynamic resizing and efficient memory management compared to arrays. They allow you to work with collections of data in a more convenient way, while also providing the ability to share underlying data between multiple slices. Understanding how slices work and how they differ from arrays is essential for effectively using them in Go programming.

package main

import (
	"fmt"
	"slices"
)

func main() {
	// uninitialized slice
	var nums []int
	fmt.Println(nums == nil) // true

	var nums2 = make([]int, 2, 5) // length: 2, capacity: 5

	// capacity -> the total number of elements that can be stored in the underlying array before it needs to be resized
	fmt.Println(cap(nums2)) // 5
	fmt.Println(nums2)

	// append elements to slice -> if the length exceeds the capacity, Go automatically allocates a new underlying array with a larger capacity and copies the existing elements to it
	nums2 = append(nums2, 3, 4, 5) // length: 5, capacity: 5
	fmt.Println(nums2)
	fmt.Println(cap(nums2))

	// Copy elements from one slice to another
	var nums3 = make([]int, 0, 5)
	nums3 = append(nums3, 2)
	var nums4 = make([]int, len(nums3))

	fmt.Println(nums3)
	
	copy(nums4, nums3) // copy elements from nums3 to nums4
	fmt.Println(nums3, nums4)

	// Slicing operations
	nums5 := []int{1, 2, 3, 4, 5}
	subSlice := nums5[:3] // Subslice -> creates a new slice that references the first three elements of nums5
	fmt.Println(subSlice) // [1 2 3]

	// Slices package
	num1 := []int{1, 2, 3}
	num2 := []int{1, 2, 4}
	// Check if num1 == num2 or not. 
	fmt.Println(slices.Equal(num1, num2)) // false

	// 2D Slices
	var nums6 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums6)
}