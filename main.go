package main

import (
	"fmt"
)

//esto es un slice, aún no entiendo este tipo de dato
var testSlice = []int{5, 6, 7, 8, 9}


func main() {

	//a continuación un bucle for comun en go
	for i := 0; i < 10; i++ {
		fmt.Println(i + 1)
	}

	//este for es muy parecido al for each de javascript, no necesita saber el lenght de el array
	//y nos devuelve también el index
	for index, value := range testSlice {
		fmt.Printf("el index es %v y el valor es %v \n", index, value)
	}
}
