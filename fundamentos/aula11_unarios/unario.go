package main

import "fmt"

func main() {
	x, y := 1, 2

	fmt.Println(x, "e o ", y)

	x++
	y--

	fmt.Println(x, "e o ", y)

	//fmt.println(++x == y--) não pode

}
