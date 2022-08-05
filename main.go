package main

import (
	"fmt"
	"runtime"
)

const (
	existError      bool   = true
	counter         int8   = 4
	operativeSystem string = runtime.GOOS
	age             int8   = 50
)

func main() {

	//este es un if else normal
	if existError == false {
		fmt.Println("no hay error")
	} else {
		fmt.Println("si hay error")
	}
	//este es un switch con una variable para hacer comparación
	switch counter {
	case 0:
		fmt.Println("the counter is cero")
	case 1:
		fmt.Println("the counter is uno")
	case 2:
		fmt.Println("the counter is dos")
		//los cases pueden tener uno o mas valores para comparar como a continuación
	case 3, 4, 5:
		fmt.Println("the counter is more than 2")
	default:
		fmt.Println("no es ninguno de los valores esperados")

	}
	//este es un switch que no tiene una variable para hacer comparación
	//la comparación se puede hacer en cada case individual lo cual puede ayudar a simplificar
	//anidaciones feas de if else
	switch {
	case operativeSystem == "Linux":
		fmt.Println("tu sistema operativo es linux")
	case operativeSystem == "darwin":
		fmt.Println("tu sistema operativo es darwin")
	default:
		fmt.Println("tu sistema operativo es otro")
	}
	//la palabra fallthrough permite que el código siga haciendo validaciones
	//para ver en que otro case puede haber una coincidencia. por ejemplo
	//en el siguiente switch se imprimen 2 valores
	switch {
	case age > 40:
		fmt.Println("ya estas como viejo")
		fallthrough
	case age == 50:
		fmt.Println("tienes 50 años")
	default:
		fmt.Println("no se que edad tienes")
	}

}
