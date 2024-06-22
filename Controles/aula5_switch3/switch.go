package main

import (
	"fmt"
	"time"
)

func checagem(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"

	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "oq é isso"
	}
}

func main() {
	fmt.Println(checagem(2.3))
	fmt.Println(checagem(1))
	fmt.Println(checagem("hellooo"))
	fmt.Println(checagem(func() {}))
	fmt.Println(checagem(time.Now()))
}
