package main

import (
	"fmt"
)

//byte: representa un caracter (una letra) con un valor numérico de tipo uint8
//las letras deben encerrarse en comillas simple cuando es una sola letra


//rune: representa un character especial(unicode) con un valor de tipo int32

func main() {

	//definiendo un byte, convierte la letra b en el número 98
	var a byte = byte('b')
	//imprimiento el valor del byte
	fmt.Println(a)

	//definiendo un arreglo de bytes
	var b []byte = []byte("Hola soy un string")
	//imprimiendo el valor del arreglo de bytes
	fmt.Println(b)

	//convirtiendo el arreglo de bytes a string
	fmt.Println(string(b))

	//creando una variable de tipo rune
	var c rune = '~'
	//imprimiendo una variable de tipo rune
	fmt.Printf("%c tiene un valor unicode de %U\n", c, c)
	//imprimiendo el valor de c como string
	fmt.Println(string(c))
}
