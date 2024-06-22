package main

import (
	"fmt"
)

func main() {
	//array n é possivel modificar as possições ou mudar o tipo

	var notas [3]float64
	//não é necessario incializar todos os elementos do array ja começam com 0
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 2.3, 10.0
	fmt.Println(notas)

	soma := 0.0

	for i := 0; i < len(notas); i++ {
		soma += notas[i]

	}
	fmt.Println("Soma", soma)
	//o len o tamanho do array é automaticamente int, ent mude isso para float
	med := soma / float64(len(notas))

	fmt.Println("media: ", med)

}
