package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("String:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("=<", 3 <= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println(d1)
	fmt.Println(d2)

	fmt.Println(d1 == d2)
	fmt.Println(d1.Equal(d2))

	type pessoa struct {
		nome string
	}

	p1 := pessoa{"joe"}

	fmt.Println(p1)

}
