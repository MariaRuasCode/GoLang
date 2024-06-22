package main

import "fmt"

func main() {

	//map incializada sem declaração de "objetos"
	teste1 := make(map[int]string)
	//nome := make(map[tipo da chave] tipo de entrada)

	//colocando variaveis
	teste1[1] = "bom dia"
	teste1[2] = "boa tarde"
	teste1[3] = "boa noite"

	fmt.Println(teste1)

	//map com objetos inicializados

	//nome := map[tipo chave] tipo do objeto {
	// chave: objeto,}
	
	teste2 := map[int]string{
		1: "bom dia",
		2: "boa tarde",
		3: "boa noite"}

	fmt.Println(teste2)

}
