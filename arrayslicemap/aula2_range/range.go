package main

import "fmt"

func main() {

	numeros := [...]int{1, 2, 3, 4, 5} //aqui o compilador é quem conta quantos elementos que tem e define o tamanho
	//TEM que ter 3 pontinho, se não tiver ent vira um slice (n sei oq isso significa ainda)

	for i, numero := range numeros {

		fmt.Printf("%v) %v\n", i+1, numero)
		//ta pelo oq eu entendi essa variavel ela só serve pra contar o tamanho de repetição, baseado no range dos elementos
	}

	//dito isso podemos ignora-la usando "_"

	for _, num := range numeros {
		fmt.Println(num)
	}

	//e se eu esquecer o negocio vai imprimir o indice

	for num2 := range numeros {
		fmt.Print(" ", num2)

		//0, 1 , 2 , 3 , 4
		//esta é a saida, mas nosso array n tem 0 ent este é o index
	}

}
