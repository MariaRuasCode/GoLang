package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroaleatorio() int {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}

func main() {
	if i := numeroaleatorio(); i > 5 { //inicializar variavel dentro daqui a torna ela privada para só aq
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu :(")
	}
}
