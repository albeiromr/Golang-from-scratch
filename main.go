package main

import (
	"fmt"
)


func main() {

	//la función fmt.Sprintf es el equivalente a los templates strings en javascript
	//retorna el string ya formateado y se le puede asignar a una variable, no imprime
	// nada en consola

	iva := 19
	grados := 32

	fraseImpuesto := fmt.Sprintf("El iva en colombia es %v%%", iva)
	fmt.Println(fraseImpuesto)

	fraseGrados := fmt.Sprintf("la temperatura en mariquita es %v° centigrados", grados)
	fmt.Println(fraseGrados)


	//-----------------------------------------------------------------------------------------------------------------------------------------
	
	//la función fmt.Printf imprime los valores según el formato o la transformación que 
	//le queramos dar, los valores se asignan de izquierda a derecha el primer format verb
	//corresponde a la primera variable declarada a la derecha y así susesivamente

	//cada tipo de dato tiene format verbs distintos que se pueden consultar aquí https://pkg.go.dev/fmt
	//también se pueden juntar dos verbos o mas como en el caso de la variable porcentaje

	porcentajeDecimal := 5632.5632
	porcentaje := 56
	texto := "Hola a todos"
	booleano := false

	fmt.Printf("el decimal es %2.2f%% el texto es %s el booleano es %t el porcentaje es %d%% ", porcentajeDecimal, texto, booleano, porcentaje)
}
