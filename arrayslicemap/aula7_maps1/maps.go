package main

import "fmt"

func main() {
	//var aprovados map[int]string

	aprovados := make(map[int]string)

	aprovados[12345454565] = "Maria"
	aprovados[12345453265] = "Pedro"
	aprovados[12345222565] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {

		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345454565])

	delete(aprovados, 12345454565)
	fmt.Println(aprovados[12345454565], "excluido")

}
