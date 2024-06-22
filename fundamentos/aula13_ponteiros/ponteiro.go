package main

import "fmt"

func main() {
	i := 1
	//go n tem aritmetica de ponteiro

	var p *int = nil //nill é null

	p = &i //pegando o endereço da variavel i

	*p++ //  * se torna a referencia

	fmt.Println(p, *p, i, &i)
}
