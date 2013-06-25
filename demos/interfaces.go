// Interfaces

package main

import "fmt"
import "strconv"

type intFactors struct {
	factorA, factorB int
}

type stringFactors struct {
	factorA          int
	stringToMultiply string
}

func (facts *intFactors) multiply() string {
	product := facts.factorA * facts.factorB
	return strconv.Itoa(product)
}

func (facts *stringFactors) multiply() string {
	newString := ""
	for i := 0; i < facts.factorA; i++ {
		newString = newString + facts.stringToMultiply
	}
	return newString
}

type Multiplier interface {
	multiply() string
}

func printMult(value Multiplier) {
	fmt.Println(value.multiply())
}

func main() {

	myFacts := intFactors{4, 5}
	myString := stringFactors{4, "Banana"}

	printMult(&myString)
	printMult(&myFacts)

}
