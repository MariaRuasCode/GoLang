package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	//se for desse jeito vai dá erro pois x é float, e y e int:
	//fmt.Println (x / y)

	//mas se transformarmos o int em float
	//agora vai
	fmt.Println(x / float64(y))

	//(podemos transformar o double em int tbm)
	fmt.Println(int(x) / y)
	//mas o valor sera arredondado

	nota := 2.5
	fmt.Println(int(nota))

	//cuidado...
	fmt.Println("teste " + string(97))
	//lembra da rune? ele retorna o valor q estiver no 97, que no caso aqui é a

	//como converter? tem um jeito mais bonito na outra aula mas tbm tem esse
	fmt.Println("int to string: " + strconv.Itoa(97))
	//Tbm tem o inverso, int para string

	num, _ := strconv.Atoi("123")
	//este _ seria como tratar um erro nesta conversão, mas como estamos ignorando

	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")

	//1 tbm considerado true

	if b {
		fmt.Println("verdadeiro")
	}
}
