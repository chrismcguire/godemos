// Structs

// No objects
// Structs: collections of fields
// Like c structs

package main

import "fmt"

type factors struct {
	factorA, factorB int
}

func multiplyFactors(facts factors) int {
	return facts.factorA * facts.factorB
}

func main() {

	myFactors := factors{4, 5}
	fmt.Println(multiplyFactors(myFactors))

}
