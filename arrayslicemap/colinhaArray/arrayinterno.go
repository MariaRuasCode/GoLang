package main

import "fmt"

func main() {

	//teste de arrays

	//aray vazio
	var teste1 []int

	//array iniciado
	var teste2 [2]int

	//array com elementos dentro
	var teste3 [3]int
	teste3[0], teste3[1], teste3[2] = 1, 2, 3

	//array inciando com elementos e tamanho indefinido
	teste4 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(teste1)
	fmt.Println(teste2)
	fmt.Println(teste3)
	fmt.Println(teste4)

}
