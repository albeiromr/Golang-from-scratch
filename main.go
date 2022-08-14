package main

import (
	"fmt"
)

//Mapas

func main() {

	//creando un mapa e inicializandolo
	aMap := map[string]int{
		"age": 15,
		"id":  128,
	}

	//creando un mapa sin inicializar y luego agregandole información
	bMap := map[string]int{}
	bMap["age"] = 56

	fmt.Println(aMap)
	fmt.Println(bMap)

	//------------------------------------------------------------------------------------------------------

	//no se pueden agregar valores a un mapa que es nil, a continuación un ejemplo
	nilMap := map[string]int{}
	nilMap["test"] = 12
	fmt.Println("nilMap no es nil: ", nilMap)
	//volviendo nilMap a tipo nil
	nilMap = nil
	fmt.Println("nilMap es nil: ", nilMap)
	//nilMap["test2"] = 12 // esta línea hace que el programa colapse por que no se pueden asignar valores
	// a un mapa que es nil
	//en el dia a dia si una función recibe un mapa como parámetro primero hay que checar si este no es nil
	//para poder comenzar a trabajar con el

	//------------------------------------------------------------------------------------------------------

	//creando un mapa con make
	cMap := make(map[string]int)
	cMap["test"] = 23
	fmt.Println(cMap)

	//------------------------------------------------------------------------------------------------------

	//iterando mapas con range
	dMap := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	}

	//cuando se intera en un map Golang no itera de forma ordenada, y es a proposito, así fue hecho el lenguaje
	for key, value := range dMap {
		fmt.Printf("el valor de %v es %v\n", key, value)
	}

	//------------------------------------------------------------------------------------------------------

}
