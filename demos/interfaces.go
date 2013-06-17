package main

import "fmt"

type intProducts struct {
	prodA, prodB int
}

type stringProducts struct {
	prodA int
    stringToMultiply string
}

type Multiplier interface {
    multiply() string
}


func (prods *intProducts) multiply() string {
    product := prods.prodA * prods.prodB
    return string(product)
}

func (prods *stringProducts) multiply() string {
    newString := ""
    for i := 0; i < prods.prodA; i++ {
        newString = newString + prods.stringToMultiply
    }
    return newString
}

func printMult(value Multiplier) {
    fmt.Println(value.multiply())
}

func main() {

	myProds := intProducts{4, 5}
    myString := stringProducts{4, "Banana"}

    printMult(&myString)
    printMult(&myProds)

    //fmt.Println(myProds.multiply())

}
