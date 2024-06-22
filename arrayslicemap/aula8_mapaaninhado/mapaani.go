package main

import "fmt"

func main() {

	Setores := map[string]map[string]float32{
		"A": {
			"Alice": 23233.4,
			"Alan":  2000.32,
		},
		"B": {
			"Bernado": 405505.23,
			"Beatriz": 33344.23,
		},
		"C": {
			"Camila": 23233.43,
			"Carol":  33434.32,
		}}

	for letra, funcs := range Setores {
		fmt.Println(letra)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}

	}
}
