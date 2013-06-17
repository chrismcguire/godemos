package main

import "fmt"

type products struct {
	prodA, prodB int
}

func multiplyProducts(prods products) int {
	return prods.prodA * prods.prodB
}

func double(factor int) int {

	return factor * 2

}

func main() {

	myProds := products{4, 5}
	fmt.Println(multiplyProducts(myProds))
}
