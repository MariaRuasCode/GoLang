package main

import (
	"fmt"
	"math"
)

// se colocar um "m" na frente do math, sera possivel referenciar este pacote com m, ou seja m.pow()
func main() {
	const PI float64 = 3.15
	var raio = 2.0

	//forma reduzida, : variavel n existe inizializando
	area := PI * math.Pow(raio, 2)
	fmt.Println("A area é: ", area)

	//muito foda isso, definir um bando de variavel dentro do ()
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// e possivel atribuir varias variaveis na mesma linha
	var e, f bool = true, false
	fmt.Println(e, f)

	//olha q simples olha q bonito, e o programa ja sabe oq cada uma é
	g, h, i := 0, false, "meu deus"
	fmt.Println(g, h, i)

	samuel := 10

	fmt.Println(samuel)
}

//constant nome tipo
//se vc atribuir um valor a variavel não é necessario definir o seu tipo

//e obrigatorio o uso de variaveis, ja as constantes n tem este problema
