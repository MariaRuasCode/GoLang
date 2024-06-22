package main

import "fmt"

func main() {

	array1 := [3]int{1, 2, 3}

	var slice1 []int
	fmt.Println(array1, slice1)

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(slice1)

	slice2 := make([]int, 2)

	//nesse copy, o slice1 tem 3 elementos e o 2 tem dois, então quando eu dou copy, para o slice2
	//acaba que o slice 2 continua com 2 primeiros elementos do slice 1
	copy(slice2, slice1)
	//copy(recebe, dá)
	fmt.Println(slice1, slice2)

}
