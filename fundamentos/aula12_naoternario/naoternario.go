package main

import "fmt"

func obterresultado(nota float64) string {

	if nota >= 6 {
		return "Aprovado"

	}
	return "Reprovado"
	//return nota >= 6 N PODE
}

func main() {
	fmt.Println(obterresultado(6.2))
}
