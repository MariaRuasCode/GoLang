package main

import "fmt"

func NotaFinal(n float64) string {

	switch {
	case n >= 9:
		return "A"
	case n >= 7:
		return "B"
	case n >= 5:
		return "C"
	case n >= 3:
		return "D"
	default:
		return "E"

	}
}

func main() {
	nota := 0.0

	fmt.Println(NotaFinal(nota))
}
