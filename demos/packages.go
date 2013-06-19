// Copy the src directory into you $GOPATH for this to work.
package main

import "complexmath"

func main() {
	myProds := complexmath.IntProducts{4, 5}
    myString := complexmath.StringProducts{4, "Banana"}
    complexmath.PrintMult(&myProds)
    complexmath.PrintMult(&myString)
}

