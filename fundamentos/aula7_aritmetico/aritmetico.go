package main

import (
	"fmt"
	"math"
)

func main() {

	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("sub => ", a-b)
	fmt.Println("div => ", a/b)
	fmt.Println("mult =>", a*b)
	fmt.Println("modulo =>", a%b)

	//em binario, tbm faz

	fmt.Println("E =>", a&b)
	fmt.Println("Ou => ", a|b)
	fmt.Println("Xor => ", a^b)

	c := 3.0
	d := 2.0

	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Print("Menor => ", math.Min(float64(c), float64(d)))

}
