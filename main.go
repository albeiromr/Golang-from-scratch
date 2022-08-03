package main

import "fmt"

//así se declaran las constantes en golang
//después de declaradas son imutables
const iva int8 = 19

//también se pueden declarar conjuntos de constantes como se realiza a continuación
const (
	diasRenta  int16   = 365
	topeDinero float32 = 54
)

func main() {

	//las constantes también se pueden declarar dentro de funciones
	const salarioMensualTope float32 = 4.200

	fmt.Println("El impuestao del iva en colombia es del ", iva, "%", "deben declarar renta cada",
		diasRenta, "y el tope de dinero es ", topeDinero, "millones")

	fmt.Println("Quienes ganen mas de", salarioMensualTope, " millones deben pagar renta")
}
