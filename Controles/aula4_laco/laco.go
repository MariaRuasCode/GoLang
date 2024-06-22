package main

import (
	"fmt"
	"time"
)

func main() {

	count := 2
	for i := 0; i <= count; i++ {
		fmt.Printf("%v ", i)
	}
	y := 1
	for y <= 10 {
		fmt.Println("transformaram for em while")
		y++
	}

	for i := 1; i < 11; i++ { //i+=2 funciona tbm

		if i%2 == 0 {
			fmt.Print(" par ")

		} else {
			fmt.Print(" impar ")
		}
	}
	fmt.Println(" ")

	fmt.Println("forever and")
	for {
		//for infinito
		fmt.Println("ever...")
		time.Sleep(time.Second)
	}
}
