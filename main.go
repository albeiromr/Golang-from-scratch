package main

import (
	"fmt"
)

var (
	intList     []int8    = []int8{1, 2, 3}
	stringList  []string = []string{"One", "Two", "Three"}
	booleanList []bool   = []bool{true, false, false, true}
)

//función sin genérico solo acepta listas de strings
func printList(list []string) {
	for _, value := range list {
		fmt.Println("values is: ", value)
	}
}

//función con genéricos acepta cualquier tipo de lista
func printListGenerics[T any](list []T) {
	for _, value := range list {
		fmt.Println("values is: ", value)
	}
}

func main() {
	//llamado sin genéricos
	printList(stringList)
	//llamados con genéricos con una lista de enteros
	printListGenerics(intList)
	//llamados con genéricos con una lista de booleanos
	printListGenerics(booleanList)
}
