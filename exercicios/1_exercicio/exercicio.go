package main

import (
	"fmt"
	"math"
)

func main() {
	lados := float64(10)

	p := 4 * lados
	a := math.Pow(lados, 2)
	d := lados * math.Sqrt(2)

	fmt.Println("Valor do Lado: ", lados)
	fmt.Println("Perimetro: ", p)
	fmt.Println("Area: ", a)
	fmt.Println("Diagonal: ", d)

}
