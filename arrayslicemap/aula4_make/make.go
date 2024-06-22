package main

import "fmt"

func main() {
	s := make([]int, 10)

	s[9] = 12

	fmt.Println(s)
	//criei um slice com 10 elementos porem criei um slice com mais elementos, mas estou refereciando somente 10, ou seja:
	//seu tamanho é 10, mas sua capacidade é 20
	s = make([]int, 10, 20)

	fmt.Println(len(s), cap(s))

	fmt.Println(s)
	//append atribui novos valores para o slice, o array continua o mesmo
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	//agora o slice esta na capacidade maxima porem se adicionar mais 1 valor

	s = append(s, 1)
	fmt.Println(len(s), cap(s))
	//21 40
	//o slice dobra de tamanho
}
