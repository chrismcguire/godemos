// Functions

// Everything pass by value, but with pointers for reference

package main

import "fmt"

func multiply(factorA, factorB int) int {

	return factorA * factorB

}

func incrementValue(value *int) {
// Increment a value and return nothing.
    *value = *value + 1

}

func main() {

	factor := 5
	fmt.Println(multiply(factor, 7))

//    incrementValue(&factor)
//    fmt.Println(factor)
}
