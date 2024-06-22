package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)

	//reflect type of indica qual tipo de variavel esta dentro do ()
	fmt.Println("literal inteiro é", reflect.TypeOf(34000))

	//sem sinal (só positios) uint8 uint16 uint32 uint64
	var b byte = 255

	fmt.Println("8 uint: ", reflect.TypeOf(b))

	valor := math.MaxInt64

	fmt.Println("maior valor int64: ", valor)
	fmt.Println("seu tipo: ", reflect.TypeOf(valor))

	//há um mapeamento na tabela unicode dentro de int32, isso indica runes, que podem tbm ser coisas como !
	//por isso quando vericamos seu tipo, teremos de volta int32
	var runa rune = 'a'

	fmt.Println("uma rune é: ", reflect.TypeOf(runa))
	//e seu valor é refente a onde esta a letra "a" dentro do unicode
	fmt.Println(runa)

	//numeros floats

	var x float32 = 34.32
	fmt.Println("O tipo x é: ", reflect.TypeOf(x))
	fmt.Println("tipo literal sempre sera 64: ", reflect.TypeOf(34.32))
	//se não especificar que a variavel é 32, ent ela sempre retorna como 64

	//boolean
	bo := true
	fmt.Println("o tipo bo é:", reflect.TypeOf(bo))
	fmt.Println(!bo)
	fmt.Println("o tipo -bo é: ", reflect.TypeOf(!bo))

	//string

	s1 := "opaaaaa"

	fmt.Println(s1 + "ssosococcoorrr")
	fmt.Println("o tamanho da string: ", len(s1))

	//char tecnicamente n existe, ele é considerado um rune

	char := 'a'

	fmt.Println("o tipo char:", reflect.TypeOf(char))
}
