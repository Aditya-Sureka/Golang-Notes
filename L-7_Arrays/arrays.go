// arrays in Go are fixed size, meaning that once you declare an array with a certain length, it cannot be resized. The length of the array is determined at compile time and cannot be changed during runtime. This is different from slices, which are more flexible and can grow or shrink in size as needed.
// In Go, arrays are value types, which means that when you assign an array to another variable or pass it to a function, a copy of the entire array is made. This can lead to performance issues if the array is large, as it requires copying all the elements. Slices, on the other hand, are reference types and do not require copying the underlying array when assigned or passed to functions, making them more efficient for handling collections of data.
// In Go, arrays are indexed starting from 0, which means that the first element of the array is accessed using index 0, the second element with index 1, and so on. The last element of the array can be accessed using the index equal to the length of the array minus one. For example, if you have an array of length 5, the valid indices for accessing its elements would be 0, 1, 2, 3, and 4.
// In Go, arrays can be multi-dimensional, meaning that you can have arrays of arrays. For example, you can declare a two-dimensional array (a matrix) using the syntax [rows][columns]type. This allows you to create more complex data structures, such as tables or grids. You can access elements in a multi-dimensional array using multiple indices, where the first index corresponds to the row and the second index corresponds to the column.
// In Go, arrays can be initialized in several ways. You can declare an array and then assign values to its elements individually, as shown in the previous example. Alternatively, you can initialize an array at the time of declaration using a composite literal, which allows you to specify the values for all elements in a concise manner. For example, you can declare and initialize an array of integers like this: nums := [5]int{1, 2, 3, 4, 5}. This creates an array of length 5 with the specified values. You can also omit the length when using a composite literal, and Go will infer it based on the number of values provided: nums := [...]int{1, 2, 3, 4, 5}.
// In Go, arrays can be passed to functions as parameters. When you pass an array to
// a function, a copy of the entire array is made, which can lead to performance issues if the array is large. To avoid this, you can pass a pointer to the array instead. This allows the function to access and modify the original array without creating a copy. For example, you can define a function that takes a pointer to an array like this: func modifyArray(arr *[5]int) { /* function body */ }. When calling this function, you would pass the address of the array using the & operator: modifyArray(&nums). This way, any changes made to the array within the function will affect the original array outside the function.
// In Go, arrays can be compared for equality using the == operator. Two arrays are considered equal if they have the same length and all corresponding elements are equal. For example, if you have two arrays of integers, arr1 and arr2, you can compare them like this: if arr1 == arr2 { /* arrays are equal */ }. However, it's important to note that arrays of different types or lengths cannot be compared directly and will result in a compile-time error. Additionally, when comparing arrays, the entire array is compared element by element, which can be inefficient for large arrays. In such cases, it may be more efficient to compare specific elements or use slices instead of arrays.
// In Go, arrays can be used in various ways, such as storing collections of data, implementing data structures, and performing operations on multiple values. However, due to their fixed size and value semantics, arrays are often less flexible and less efficient than slices for most use cases. Slices provide a more convenient and efficient way to work with collections of data in Go, as they allow for dynamic resizing and do not require copying the entire array when passed to functions or assigned to new variables. Therefore, while arrays can be useful in certain situations where a fixed-size collection is needed, slices are generally the preferred choice for working with collections of data in Go.
// In summary, arrays in Go are fixed-size collections of elements that are indexed starting from 0. They are value types, meaning that they are copied when assigned or passed to functions. Arrays can be multi-dimensional and can be initialized in various ways. However, due to their limitations, slices are often preferred for working with collections of data in Go.

package main

import "fmt"

func main() {
	// length of array
	var nums [8]int
	fmt.Println(len(nums))

	// assign value to array
	nums[0] = 1
	fmt.Println(nums[0])

	// In single line
	nums2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums2)

	// multi-dimensional array (2D Array)
	matrix := [2][2]int{{3, 4},
		                {5, 6}}
	fmt.Println(matrix)
}
