// Methods


package main

import "fmt"

type factors struct {
	factorA, factorB int
}

// Define a receiver for the function
func (facts *factors) multiply() int {
// Don't need the pointer to products, but use it most of the time (don't have to copy)

    return facts.factorA * facts.factorB

}

func main() {

	myFactors := factors{4, 5}
    // Call multiply on the type
    fmt.Println(myFactors.multiply())

}
