package main

import "fmt"

func comprar(trab1, trab2 bool) (bool, bool, bool) {

	comprartv50 := trab1 && trab2
	comprartv32 := trab1 != trab2 // não existe "ou" exclusivo no go lang :(
	comprarsorvete := trab1 || trab2

	return comprartv50, comprartv32, comprarsorvete
}

func main() {
	tv50, tv32, sorvete := comprar(true, true)

	fmt.Printf("Tv50: %v, tv32: %v, Sorvete: %v, Saudavel: %v ", tv50, tv32, sorvete, !sorvete)
}
