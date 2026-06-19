// Maps in Go are unordered collections of key-value pairs.
// They are also known as hash maps or dictionaries in other programming languages.
// Maps provide a way to associate values with keys, allowing for efficient retrieval of values based on their corresponding keys.
// In Go, maps are created using the built-in make function or by using a map literal.
// The syntax for declaring a map is as follows: var myMap map[keyType]valueType.
// For example, you can declare a map that associates strings with integers like this: var ageMap map[string]int.
// This creates an empty map that can hold string keys and integer values.
// You can also initialize a map with values at the time of declaration using a composite literal, like this: ageMap := map[string]int{"Alice": 30, "Bob": 25}.
// This creates a map with two key-value pairs, where "Alice" is associated with 30 and "Bob" is associated with 25.
// Maps in Go are implemented as hash tables, which provide efficient lookup and insertion operations.
// When you access a value in a map using its key, Go computes the hash of the key to determine the location of the corresponding value in the underlying data structure.
// This allows for fast retrieval of values based on their keys, even in large maps.
// However, it's important to note that the order of key-value pairs in a map is not guaranteed, as maps are unordered collections.
// If you need to maintain a specific order of key-value pairs, you may need to use a different data structure, such as a slice of structs or a sorted map implementation.
// Maps in Go also support various operations, such as adding, updating, and deleting key-value pairs.
// You can add or update a key-value pair in a map by simply assigning a value to a key, like this: ageMap["Charlie"] = 35.
// This adds a new key "Charlie" with the value 35 to the map. If the key already exists, it updates the value associated with that key.
// You can delete a key-value pair from a map using the built-in delete function, like this: delete(ageMap, "Bob").
// This removes the key "Bob" and its associated value from the map.
// In summary, maps in Go are powerful data structures that allow you to associate values with keys for efficient retrieval.
// They are implemented as hash tables and provide various operations for adding, updating, and deleting key-value pairs.
// Understanding how maps work and how to use them effectively is essential for working with collections of data in Go programming.

package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// get an element
	fmt.Println(m["name"], m["area"])

	/* IF KEY,VALUE DOESN'T EXIST IN MAP, IT RETURNS 0 VALUE */
	fmt.Println(m["nonexistent"]) // returns empty string / 0

	// delete an element
	delete(m, "area") // outputs map[name:golang]

	fmt.Println(m) // prints complete map

	clear(m) // clears the map

	// if elements already present
	mp := map[string]int{"price":40, "phones": 3};
	fmt.Println(mp) // prints map[phones:3 price:40]

	// check if key exists in map
	v, ok := mp["price"] // ok is true if key exists, false otherwise
	fmt.Println(v)
	if ok {
		fmt.Println("all okay")
	} else {
		fmt.Println("not okay")
	}

	// check if two maps are equal
	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 3}
	fmt.Println(maps.Equal(m1, m2))

}