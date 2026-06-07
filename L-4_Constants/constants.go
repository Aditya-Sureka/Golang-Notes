/* Code for Declaring Constants in Go */

package main

import "fmt"

func main() {
	const name = "golang"

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}