package main

import (
	"fmt"
)

func main() {

	//Escape Characters or scape sequences

    //se usan para insertar caracteres que son ilegales en un string, son como las entidades en html
	//estan presentes en muchos lenguajes de programación, a continuación los que encontré relevantes



	// (\") Double quote -----------
	fmt.Println("scape sequence para crear \"comillas dobles\" dentro de un string")
	//output: esta es una "línea" con un scape character

	// (\\) back-slash -----------
	fmt.Println("scape sequence para crear un back-slash \\ dentro de un string")
	//output: cape sequence para crear un back-slash \ dentro de un string

	// (\b) retroceder un espacio -----------
	fmt.Println("el porcentaje es del", 5, "\b%")
	//output: el porcentaje es del 5%

	// (\n) salto de línea -----------
	fmt.Println("esta es la primera línea, \n esta es la segunda línea")
	//esta es la primera línea,
	//esta es la segunda línea

}
