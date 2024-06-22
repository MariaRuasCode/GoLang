package main

import "fmt"

func main() {

	funcEsalario := map[string]float64{
		"Jose":     1212.34,
		"Gabriel":  1200.4,
		"Fernanda": 13456.4}

	funcEsalario["Leandro"] = 1350.0

	for nome, salario := range funcEsalario {
		fmt.Println(nome, salario)
	}

	fmt.Println("Ignorando o nome")
	for _, salario := range funcEsalario {
		fmt.Println(salario)
	}

}
