package main

import "fmt"

func main() {
	a1 := [...]int{1, 2, 3, 4, 5, 6}

	//nome := [tamanho] tipo {elementos}

	s1 := a1[0:3]
	s2 := a1[3:6]

	fmt.Println(a1, s1, s2)

	//os slice aponta para o mesmo array interno!
	s3 := make([]int, 10, 20)
	s4 := append(s3, 1, 2, 3)

	s3[0] = 7

	fmt.Println(s3, s4)
}
