package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3
	fmt.Println(i)
	i += 3
	fmt.Println(i)
	i -= 3
	fmt.Println(i)
	i /= 2
	fmt.Println(i)
	i *= 3
	fmt.Println(i)
	i %= 3
	fmt.Println(i)

	x, y := 1, 2
	fmt.Printf("num x = %v, num y = %v", x, y)
	x, y = y, x
	fmt.Printf("num x = %v, num y = %v", x, y)

}
