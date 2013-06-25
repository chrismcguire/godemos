package complexmath

import "fmt"
import "strconv"

type IntProducts struct {
	ProdA, ProdB int
}

type StringProducts struct {
	ProdA int
    StringToMultiply string
}

type Multiplier interface {
    Multiply() string
}

func (prods *IntProducts) Multiply() string {
    product := prods.ProdA * prods.ProdB
    return strconv.Itoa(product)
}

func (prods *StringProducts) Multiply() string {
    newString := ""
    for i := 0; i < prods.ProdA; i++ {
        newString = newString + prods.StringToMultiply
    }
    return newString
}

func PrintMult(value Multiplier) {
    fmt.Println(value.Multiply())
}
