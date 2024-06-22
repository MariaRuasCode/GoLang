package main

import (
	"fmt"
)

func main() {
	//funciona iqual o console.write
	fmt.Print("oi")
	fmt.Print(" mesma linha")

	//funciona iqual o console.writeline
	fmt.Println(" ")
	fmt.Println("agora")
	fmt.Println("quebrou?")

	x := 4.34345
	//ele transforma o x em string
	xs := fmt.Sprint(x)
	fmt.Println("o valor de x é " + xs)

	//fazer isso sem usar sprint =
	fmt.Println("o valor de x é", x)

	//print tf é um print formatado
	fmt.Printf(" printf: O valor de x é %f", x)
	fmt.Println(" ")

	//imprime só 2 casas decimais
	fmt.Printf(" printf: O valor de x é %.2f", x)

	a, b, c, d := 1, 1.223, false, "ola"

	//tipo de variaveis tem letra diferentes
	//f são float, t são bool, s string, e d decimais

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	//o v imprimi todos da mesma maneira
	fmt.Printf("\n%v %v %.1v %v %v", a, b, b, c, d)

}
